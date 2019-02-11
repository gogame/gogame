package video

type Device interface {
	Start()
	Update()
	Exit()
}
