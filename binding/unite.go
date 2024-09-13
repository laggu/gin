package binding

type uniteBinding struct {
	bindings []Binding
}

func NewUniteBinding(bindings ...Binding) Binding {
	return &uniteBinding{
		bindings: bindings,
	}
}

func (uniteBinding) Name() string {
	return "unite"
}

func (b uniteBinding) Bind(c context, obj any) error {
	if err := b.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (b uniteBinding) mapping(c context, obj any) error {
	for _, binding := range b.bindings {
		if err := binding.mapping(c, obj); err != nil {
			return err
		}
	}
	return nil
}
