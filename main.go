package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stripe/stripe-go/v72"
	"github.com/zhashkevych/todo-app/modules/config"
	"github.com/zhashkevych/todo-app/modules/user"
	"github.com/zhashkevych/todo-app/pcg/loging"

	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)
func main() {

	fmt.Println("here its payment processing")

	stripe.Key = "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"

	logger := loging.GetLogger() //получаем логер

	logger.Info("crate router")

	router := httprouter.New() //создание роутера

	conf := config.GetConfig() //получаем конфиг


	fmt.Println("Successfully created connection to database")

	fmt.Println("Inserted 3 rows of data")

	logger.Info("registr user handler")

	handler:=user.NewHandler(logger)

	handler.Register(router)


	start(router, conf)

}


/*фуникция для подключение к локал хосту*/
func start(router* httprouter.Router,conf*config.Config)  {

	logger:= loging.GetLogger()

	logger.Info("start aplication")

	var listener net.Listener

	var listenErr error //абсолютный путь на бинарник

	if conf.Listen.Type == "sock" {

		logger.Info("detect app path")

		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))//абсолютный путь на бинарник

		if err != nil {

			logger.Fatal(err)

		}

		logger.Info("create socket")

		socketPath := path.Join(appDir, "app.sock")//склеиваем=)

		logger.Info("listen unix socket")

		listener, listenErr = net.Listen("unix", socketPath)

		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Listen.BindIP, conf.Listen.Port))
		logger.Infof("server is listening port %s:%s",conf.Listen.BindIP, conf.Listen.Port)
	}

	if listenErr != nil {
		logger.Info(listenErr)
	}



	listener,err:=net.Listen("tcp",":9001")//порт сервера

	if err!=nil{

		panic(err)

	}

	server:=&http.Server{

		Handler: router,

		/*15 в будуешем нужно посчитать и поменять */
		WriteTimeout:15 *time.Second,

		ReadTimeout: 15*time.Second,

	}

	log.Fatalln(server.Serve(listener))

}