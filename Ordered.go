package priority

type Ordered interface {
	Less(v Ordered) bool
}
