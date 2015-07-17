package angularjs

const AnnotServices = "ajs-service"

func Annotate(cls interface{}, typ string) (interface{}, error) {
	return buildParams(cls, typ)
}
