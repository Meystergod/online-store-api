package product

import pb_online_store_products "github.com/Meystergod/online-store-grpc-contracts/gen/go/online_store/products/v1"

type Server struct {
	pb_online_store_products.UnimplementedProductServiceServer
}

func NewServer(srv pb_online_store_products.UnimplementedProductServiceServer) *Server {
	return &Server{UnimplementedProductServiceServer: srv}
}
