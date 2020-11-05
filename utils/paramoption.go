package utils

type ParamOptionFunc func(ParamOption)

type ParamOption interface {
	Apply(...ParamOptionFunc)
}
