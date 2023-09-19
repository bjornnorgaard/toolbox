package gopts

func Build[OptsType any](opts OptsType, changes ...func(opts *OptsType)) OptsType {
	for _, apply := range changes {
		apply(&opts)
	}
	return opts
}
