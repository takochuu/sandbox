package architecture

import (
	"github.com/takochuu/sandbox/archtecture/postgres"
	"github.com/takochuu/sandbox/archtecture/http"
	"log"
)

func main()  {
	db, err := postgres.Open(os.getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := &postgres.UserService{DB: db}

	var h http.Handler
	h.UserService = us
}
