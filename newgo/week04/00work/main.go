package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

/*
- 考察点：map增量更新+接口+结构体方法
    - jobManager 增量更新job
    - 要求写接口，有start stop hash三个方法
    - 写两个结构体，分别实现上述结构
    - 写一个jobmanager管理，要求有增量更新
    - 远端sync
        - 本地有，远端没有，要删除
        - 本地没有，远端有，要新增
        - 本地有，远端有，不管

*/

/*
编写过程
1. 先写接口
2. 至少两个结构体实现这些接口
	- 先写结构体
	- 绑定接口中的方法
3. 有一个jobManager的结构体，有一个对应的变量，作用管理
	- activeJobs 是一个map，hold住上次的常驻的job
4. 要有一个远端sync的方法，触发增量更新
	- go ：定时去mysql中查询
*/
type jobManager struct {
	targetMtx  sync.RWMutex
	activeJobs map[string]job
}

func (jm *jobManager) getDetail() {
	ks := []string{}
	hs := []string{}
	jm.targetMtx.RLock()
	for _, v := range jm.activeJobs {
		if strings.HasPrefix(v.hash(), "k8s") {
			ks = append(ks, v.hash())
		} else {
			hs = append(hs, v.hash())
		}
	}
	jm.targetMtx.RUnlock()
	log.Printf("[k8s:%d][host:%d][k8s.detail:%s][host.detail:%s]",
		len(ks),
		len(hs),
		strings.Join(ks, ","),
		strings.Join(hs, ","),
	)
}

//增量更新
func (jm *jobManager) sync(jobs []job) {
	//待开启的
	thisNewJobs := make(map[string]job)
	//远端全量
	thisAllJobs := make(map[string]job)
	jm.targetMtx.Lock()
	//判断本次新的
	for _, j := range jobs {
		hash := j.hash()
		thisAllJobs[hash] = j
		if _, loaded := jm.activeJobs[hash]; !loaded {
			thisNewJobs[hash] = j
			jm.activeJobs[hash] = j
		}
	}
	//判断旧的并删除
	for hash, t := range jm.activeJobs {
		if _, loaded := thisAllJobs[hash]; !loaded {
			t.stop()
			delete(jm.activeJobs, hash)
		}
	}
	jm.targetMtx.Unlock()
	for _, t := range thisNewJobs {
		t.start()
	}
	jm.getDetail()
}

type job interface {
	hash() string
	start()
	stop()
}
type k8sJob struct {
	Id           int
	Name         string
	k8sNameSpace string
}

func (kj *k8sJob) hash() string {
	return "k8s" + kj.Name
}

func (kj *k8sJob) start() {
	log.Printf("[k8s.job.start][%v]", kj)
}

func (kj *k8sJob) stop() {
	log.Printf("[k8s.job.start][%v]", kj)
}

type hostJob struct {
	Id     int
	Name   string
	HostIp string
}

func (hj *hostJob) hash() string {
	return "host" + hj.Name
}

func (hj *hostJob) start() {
	log.Printf("[host.job.start][%v]", hj)
}

func (hj *hostJob) stop() {
	log.Printf("[host.job.start][%v]", hj)
}

func main() {
	jm := &jobManager{
		activeJobs: make(map[string]job),
	}
	jobs := make([]job, 0)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("k8s_job_%d", i)
		namespace := fmt.Sprintf("k8s_anemspace_%d", i)
		j := k8sJob{
			Id:           i,
			Name:         name,
			k8sNameSpace: namespace,
		}
		jobs = append(jobs, &j)
	}

	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("host_job_%d", i)
		ip := fmt.Sprintf("1.1.1.%d", i)
		j := hostJob{
			Id:     i,
			Name:   name,
			HostIp: ip,
		}
		jobs = append(jobs, &j)
	}
	log.Printf("[分配给我6个job]，部署3个k8s 3个host")
	jm.sync(jobs)
	time.Sleep(5 * time.Second)
	log.Printf("[等待5秒，下一轮分配]")

	jobs = make([]job, 0)

	for i := 1; i < 6; i++ {
		name := fmt.Sprintf("k8s_job_%d", i)
		namespace := fmt.Sprintf("namespace_%d", i)
		j := k8sJob{
			Id:           i,
			Name:         name,
			k8sNameSpace: namespace,
		}
		jobs = append(jobs, &j)
	}
	log.Printf("[分配给我5个job，部署5个k8s]")
	jm.sync(jobs)

	time.Sleep(5 * time.Second)
	log.Printf("[等待5秒，下一轮分配]")

	jobs = make([]job, 0)

	for i := 2; i < 5; i++ {
		name := fmt.Sprintf("host_job_%d", i)
		ip := fmt.Sprintf("1.1.1.%d", i)
		j := hostJob{
			Id:     i,
			Name:   name,
			HostIp: ip,
		}
		jobs = append(jobs, &j)
	}
	log.Printf("[分配给我3个job， 3个host]")
	jm.sync(jobs)
}
