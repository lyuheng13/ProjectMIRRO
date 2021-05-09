package model

type Category struct {
	Action    int `json:"Action"`
	Love      int `json:"Love"`
	Suspense  int `json:"Suspense"`
	Comedy    int `json:"Comedy"`
	Horror    int `json:"Horror"`
	Family    int `json:"Family"`
	Music     int `json:"Music"`
	Dance     int `json:"Dance"`
	Adventure int `json:"Adventure"`
	History   int `json:"History"`
	Magic     int `json:"Magic"`
	War       int `json:"War"`
	Crime     int `json:"Crime"`
	Sad       int `json:"Sad"`
	Happy     int `json:"Happy"`
	Angry     int `json:"Angry"`
	Exciting  int `json:"Exciting"`
}

type Rate struct {
	ID    string `json:"userid"`
	Type  string `json:"type"`
	Score int    `json:"score"`
	Like  bool   `json:"like"`
}
