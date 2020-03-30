package api

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver type
type Resolver struct{}

// func (r *fileResolver) MountPoint(ctx context.Context, obj *File) (*string, error) {
// 	panic("not implemented")
// }

// // File returns FileResolver implementation.
// func (r *Resolver) File() FileResolver { return &fileResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type fileResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
