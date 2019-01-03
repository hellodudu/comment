package res

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type XmlEntityClient struct {
	Id           string `xml:"id,attr"`
	Name         string `xml:"name,attr"`
	Icon         string `xml:"icon,attr"`
	Model        string `xml:"model,attr"`
	Desc         string `xml:"desc,attr"`
	Aptitude     string `xml:"aptitude,attr"`
	Scale        string `xml:"scale,attr"`
	AngrySkill   string `xml:"angry_skill,attr"`
	SortPriority string `xml:"sort_priority,attr"`
	IconType     string `xml:"icon_type,attr"`
	RareIcon     string `xml:"rare_icon,attr"`
}

type XmlLoader struct {
	EntityClientXml XmlEntityClient `xml:"entity"`
}

func NewXmlLoader() *XmlLoader {
	xmlLoader := &XmlLoader{}
	xmlFile, err := os.Open("entity_client.xml")
	if err != nil {
		log.Println(err)
	}
	defer xmlFile.Close()

	byXml, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Println(err)
	}

	if err := xml.Unmarshal(byXml, xmlLoader); err != nil {
		log.Println(err)
	}

	log.Printf("Unmarshal xml success : %+v\n", xmlLoader)

	return xmlLoader
}
