package es_log

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"
)

func NewGinLog(esHost, urls, index string) *logrus.Logger {
	log := logrus.New()

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(urls))
	if err != nil {
		log.Panicf("Elastic not reachable\n %s", err.Error())
	}
	hook, err := elogrus.NewAsyncElasticHook(client, esHost, logrus.InfoLevel, index)
	if err != nil {
		log.Panic(err)
	}
	log.Hooks.Add(hook)
	log.Infof("GinLogConfig ring in %s - %s", esHost, index)

	return log
}
