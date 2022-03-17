package api

import (
	"fmt"
	"github.com/checkpoint-restore/go-criu/v5"
	"github.com/checkpoint-restore/go-criu/v5/rpc"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"go-criu/cmd"
	response "go-criu/model"
	"go-criu/utils"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

// @Summary Freezing process
// @Description Freezing process
// @Tags ciu
// @Accept  application/json
// @Product application/json
// @Param data body response.DumpParam true "Request parameters"
// @Success 200 {object} response.Response true "Response parameters"
// @Router /dump [post]
func Dump(gin *gin.Context) {
	param := response.DumpParam{}
	if err := gin.BindJSON(&param); err != nil {
		log.Println("Dump request param is invalid, please check!", err)
		response.Fail(gin, http.StatusBadRequest, fmt.Sprint("Dump request param is invalid, please check!", err), "")
		return
	}
	c := criu.MakeCriu()
	if err := c.Prepare(); err != nil {
		log.Print("c.Prepare() error：", err)
		response.Fail(gin, http.StatusBadRequest, fmt.Sprint("c.Prepare() error：", err), "")
		return
	}
	//Check whether the storage path exists
	utils.PathUrlExists(param.Dir)
	img, err := os.Open(param.Dir)
	if err != nil {
		log.Println(err, "Can't open image dir (%s)")
		response.Fail(gin, http.StatusBadRequest, fmt.Sprint("Can't open image dir", err), "")
		return
	}
	defer img.Close()
	opts := rpc.CriuOpts{
		Pid:             proto.Int32(param.Pid),
		ImagesDirFd:     proto.Int32(int32(img.Fd())),
		LogLevel:        proto.Int32(4),
		LogFile:         proto.String("dump.log"),
		TcpEstablished:  proto.Bool(param.TcpEstablished),
		ShellJob:        proto.Bool(param.ShellJob),
		TcpSkipInFlight: proto.Bool(param.TcpSkipInFlight),
		LeaveRunning:    proto.Bool(param.LeaveRunning),
		TcpClose:        proto.Bool(param.TcpClose),
	}

	if err = c.Dump(&opts, criu.NoNotify{}); err != nil {
		if cmd.FlagNs {
			log.Printf("criu-ns Dump Pid %s imgDir %s", strconv.Itoa(int(param.Pid)), param.Dir)
			//Create a new namespace through criu-ns
			command := exec.Command("./criu-ns", "dump", "-t", strconv.Itoa(int(param.Pid)), "-D", param.Dir, "-o", "dump.log", "--tcp-established", "-j", "-v4", "--shell-job")
			err := command.Start()
			if err != nil {
				log.Printf("criu-ns Dump error ：%s", err)
				response.Fail(gin, http.StatusBadRequest, err.Error(), "")
				return
			}
			go command.Wait()
			response.Success(gin, http.StatusText(http.StatusOK), command.Process.Pid)
			return
		} else {
			log.Println("Dump failed", err)
			response.Fail(gin, http.StatusFailedDependency, err.Error(), "")
			return
		}
	}
	c.Cleanup()
	response.Success(gin, http.StatusText(http.StatusOK), "")
	return
}

// @Summary Resume frozen programs
// @Description Resume frozen programs
// @Tags ciu
// @Accept  application/json
// @Product application/json
// @Param data body response.RestoreParam true "Request parameters"
// @Success 200 {object} response.Response true "Response parameters"
// @Router /restore [post]
func Restore(gin *gin.Context) {
	param := response.RestoreParam{}
	if err := gin.BindJSON(&param); err != nil {
		log.Println("RestoreParam request param is invalid", err)
		response.Fail(gin, http.StatusBadRequest, fmt.Sprint("RestoreParam request param is invalid, please check!", err), "")
		return
	}
	c := criu.MakeCriu()
	img, err := os.Open(param.Dir)
	if err != nil {
		log.Println("Can't open image dir", err)
		response.Fail(gin, http.StatusBadRequest, fmt.Sprint("Can't open image dir error:", err), "")
		return
	}
	defer img.Close()

	opts := rpc.CriuOpts{
		LogLevel:        proto.Int32(4),
		LogFile:         proto.String("restore.log"),
		ImagesDirFd:     proto.Int32(int32(img.Fd())),
		TcpEstablished:  proto.Bool(param.TcpEstablished),
		TcpSkipInFlight: proto.Bool(param.TcpSkipInFlight),
		ShellJob:        proto.Bool(param.ShellJob),
		LeaveRunning:    proto.Bool(param.LeaveRunning),
		TcpClose:        proto.Bool(param.TcpClose),
		RstSibling:      proto.Bool(param.RstSibling),
	}

	if err = c.Restore(&opts, criu.NoNotify{}); err != nil {
		if cmd.FlagNs {
			log.Printf("criu-ns Restore -D %s", param.Dir)
			//Create a new namespace through criu-ns
			command := exec.Command("./criu-ns", "restore", "-D", param.Dir, "-o", "restore.log", "-d", "--tcp-close", "-j", "-v4", "--shell-job")
			if err := command.Start(); err != nil {
				log.Printf("criu-ns Restore error：%s", err)
				response.Fail(gin, http.StatusInternalServerError, err.Error(), "")
				return
			}
			go command.Wait()
			response.Success(gin, http.StatusText(http.StatusOK), command.Process.Pid)
			return
		} else {
			log.Println("Restore failed：", err)
			response.Fail(gin, http.StatusInternalServerError, err.Error(), "")
			return
		}
	}
	c.Cleanup()
	response.Success(gin, http.StatusText(http.StatusOK), "")
	return
}
