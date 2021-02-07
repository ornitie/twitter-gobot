package resources

type TweetsResource struct {
	baseResource *BaseResource
}

func NewTweetsResource(base *BaseResource) *TweetsResource {
	return &TweetsResource{
		baseResource: base,
	}
}
