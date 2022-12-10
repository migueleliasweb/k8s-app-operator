package secondary_reconcilers

var EnabledReconcilers = []ReconcilerWithApp{
	&Deployment{},
}
