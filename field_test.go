package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Khannedy").Info("Hello World")

	logger.WithField("name", "Eko").
		WithField("username", "Kurniawan").
		WithField("nama", "Bayu").
		Info("Hello")

}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	//menambahkan fiels atau isi pada logging panel secara banyak
	logger.WithFields(logrus.Fields{
		"username": "Eko",
		"nama":     "huhuhu",
	}).Info("hello")
}
