package trpc

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"toxotes.co/core/pkg/terror"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*terror.Err); ok {
			logx.WithContext(ctx).Errorf("[RCP_SRV_ERR] %+v", err)
			err = status.Error(codes.Code(e.GetCode()), e.GetMsg())
		} else {
			logx.WithContext(ctx).Errorf("[RCP_SRV_ERR] %+v", err)
		}
	}
	return resp, err
}
