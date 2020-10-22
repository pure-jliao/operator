package constants

const (
	// OperatorPrefix is a prefix to use for annotations added by the operator
	OperatorPrefix = "operator.libopenstorage.org"
	// LabelKeyClusterName is the name of the label key for the cluster name
	LabelKeyClusterName = OperatorPrefix + "/name"
	// LabelKeyDriverName is the name of the label key for the storage driver for the cluster
	LabelKeyDriverName = OperatorPrefix + "/driver"
	// LabelKeyStoragePod is the name of the key for the label on the pod that indicates it's on a storage node
	LabelKeyStoragePod = "storage"
	// LabelKeyKVDBPod is the name of the key for the label on the pod that indicates it's on a KVDB node
	LabelKeyKVDBPod = "kvdb"
	// LabelValueTrue is the constant for a "true" label value
	LabelValueTrue = "true"
	// AnnotationDisableStorage annotation to disable the storage pods from running (default: false)
	AnnotationDisableStorage = OperatorPrefix + "/disable-storage"
	// AnnotationReconcileObject annotation to toggle reconciliation of operator created objects
	AnnotationReconcileObject = OperatorPrefix + "/reconcile"
	// AnnotationClusterAPIMachine is the annotation key name for the name of the
	// machine that's backing the k8s node
	AnnotationClusterAPIMachine = "cluster.k8s.io/machine"
	// AnnotationCordonedRestartDelay is the annotation key name for the duration
	// (in seconds) to wait before restarting the storage pods
	AnnotationCordonedRestartDelay = OperatorPrefix + "/cordoned-restart-delay-secs"
)
