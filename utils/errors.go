package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var TestError = status.Error(codes.Internal, "test error")
var NotFoundError = status.Error(codes.NotFound, "")
var UsageIntError = status.Error(codes.InvalidArgument, "use: 1/0/-1 as login to get good/bad/error test responses")
var UsageStringError = status.Error(codes.InvalidArgument, "use: good/bad/error as login to get appropriate test responses")
