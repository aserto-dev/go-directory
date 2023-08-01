package model

import (
	"context"
	"io"
	"net/http"

	dms3 "github.com/aserto-dev/go-directory/aserto/directory/model/v3"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const MaxChunkSizeBytes int = 64 * 1024

func RegisterModelStreamHandlersFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterModelStreamHandlerClient(ctx, mux, dms3.NewModelClient(conn))
}

func RegisterModelStreamHandlerClient(ctx context.Context, mux *runtime.ServeMux, client dms3.ModelClient) error {
	if err := mux.HandlePath(
		"GET",
		"/api/v3/directory/manifest/{name}",
		getManifestHandler(mux, client),
	); err != nil {
		return errors.Wrap(err, "failed to register GetManifest handler")
	}

	if err := mux.HandlePath(
		"POST",
		"/api/v3/directory/manifest/{name}/{version}",
		setManifestHandler(mux, client),
	); err != nil {
		return errors.Wrap(err, "failed to register SetManifest handler")
	}

	return nil
}

func getManifestHandler(mux *runtime.ServeMux, client dms3.ModelClient) runtime.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		ctx, err := runtime.AnnotateContext(
			req.Context(),
			mux,
			req,
			"/aserto.directory.model.v3.Model/GetManifest",
			runtime.WithHTTPPathPattern("/api/v3/directory/manifest/{name}"),
		)
		if err != nil {
			runtime.HTTPError(req.Context(), mux, outboundMarshaler, w, req, err)
			return
		}

		stream, err := client.GetManifest(ctx, &dms3.GetManifestRequest{
			Name:    pathParams["name"],
			Version: req.URL.Query().Get("version"),
		})
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		w.Header().Set("Content-Type", "application/yaml")
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}

			if md := msg.GetMetadata(); md != nil {
				w.Header().Set("X-Created-At", md.CreatedAt.AsTime().Format(http.TimeFormat))
				w.Header().Set("X-Updated-At", md.UpdatedAt.AsTime().Format(http.TimeFormat))
				w.Header().Set("ETag", md.Etag)
			}

			if body := msg.GetBody(); body != nil {
				if _, err := w.Write(body.Data); err != nil {
					runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
					return
				}
			}
		}
	}
}

func setManifestHandler(mux *runtime.ServeMux, client dms3.ModelClient) runtime.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		ctx, err := runtime.AnnotateContext(
			req.Context(),
			mux,
			req,
			"/aserto.directory.model.v3.Model/SetManifest",
			runtime.WithHTTPPathPattern("/api/v3/directory/manifest/{name}/{version}"),
		)
		if err != nil {
			runtime.HTTPError(req.Context(), mux, outboundMarshaler, w, req, err)
			return
		}
		stream, err := client.SetManifest(ctx)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		if err := stream.Send(&dms3.SetManifestRequest{
			Msg: &dms3.SetManifestRequest_Metadata{Metadata: &dms3.Metadata{
				Name:    pathParams["name"],
				Version: pathParams["version"],
			}},
		}); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		reader := req.Body
		defer reader.Close()
		buf := make([]byte, 1024)
		for {
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}

			if n > 0 {
				if err := stream.Send(&dms3.SetManifestRequest{
					Msg: &dms3.SetManifestRequest_Body{Body: &dms3.Body{Data: buf[:n]}},
				}); err != nil {
					runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
					return
				}
			}

			if err == io.EOF {
				break
			}
		}

		if _, err := stream.CloseAndRecv(); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
	}
}
