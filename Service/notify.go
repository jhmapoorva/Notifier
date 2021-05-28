package Service

import "context"

func (s Server) Notify(ctx context.Context, request *NotifyRequest) (*NotifyResponse, error) {
	//TODO: This error does not make sense as there might not be any key for that specific district therefore should be ignored, same with telegram need to check
	telegram()
	err := tweet()
	if err != nil {
		return nil, err
	}
	return &NotifyResponse{Success: true}, nil
}
