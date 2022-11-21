package infrastructure

// import (
// 	"github.com/gorilla/mux"
// 	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
// )

// type ControllHandler struct {
// 	Common *interfaces.CommonController
// }

// func NewServer(h db.SqlHandler) (handler *mux.Router) {
// 	// Handler
// 	ch := &ControllHandler{
// 		Common: interfaces.NewController(h), // Controller増えるごとに追加
// 		// Admin: adminController // 初期化されたコントローラー追加
// 	}
// 	handler = NewRouter(ch)
// 	return
// }
