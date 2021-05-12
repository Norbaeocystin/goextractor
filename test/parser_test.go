package test

import (
	"goextractor/pkg/goextractor"
	"testing"
)

func TestClean(t *testing.T){
	domain := goextractor.Clean("https://www.topky.sk/cl/10/2113533/Vahate--Tu-je-5-DOVODOV--preco-neodkladat-ockovanie-proti-covidu")
	if domain != "topky.sk"{
		t.Errorf("Clean does not work")
	}
}

func TestCleanLink(t *testing.T){
	url := goextractor.CleanLink("https://www.topky.sk?bla=fha")
	if url != "https://www.topky.sk"{
		t.Errorf("Clean does not work")
	}
	url = goextractor.CleanLink("https://www.topky.sk#fha")
	if url != "https://www.topky.sk"{
		t.Errorf("Clean does not work")
	}
}