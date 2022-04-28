package epub

import "encoding/xml"

type Package struct {
	XMLName          xml.Name
	Version          string    `xml:"version,attr,omitempty"`
	UniqueIdentifier string    `xml:"unique-identifier,attr,omitempty"`
	ID               string    `xml:"id,attr,omitempty"`
	Prefix           string    `xml:"prefix,attr,omitempty"`
	Lang             string    `xml:"lang,attr,omitempty"`
	Dir              string    `xml:"dir,attr,omitempty"`
	Opf              string    `xml:"opf,attr,omitempty"`
	Ns               string    `xml:"ns,attr,omitempty"`
	Metadata         *Metadata `xml:"metadata,omitempty"`
	Manifest         *struct {
		ID   string `xml:"id,attr,omitempty"`
		Item *[]struct {
			Text         string `xml:",chardata"`
			ID           string `xml:"id,attr,omitempty"`
			Href         string `xml:"href,attr,omitempty"`
			MediaType    string `xml:"media-type,attr,omitempty"`
			Fallback     string `xml:"fallback,attr,omitempty"`
			MediaOverlay string `xml:"media-overlay,attr,omitempty"`
			Properties   string `xml:"properties,attr,omitempty"`
		} `xml:"item,omitempty"`
	} `xml:"manifest,omitempty"`
	Spine *struct {
		ID                       string `xml:"id,attr,omitempty"`
		Toc                      string `xml:"toc,attr,omitempty"`
		PageProgressionDirection string `xml:"page-progression-direction,attr,omitempty"`
		Itemref                  *[]struct {
			Text       string `xml:",chardata"`
			Idref      string `xml:"idref,attr,omitempty"`
			Linear     string `xml:"linear,attr,omitempty"`
			ID         string `xml:"id,attr,omitempty"`
			Properties string `xml:"properties,attr,omitempty"`
		} `xml:"itemref,omitempty"`
	} `xml:"spine,omitempty"`
	Guide *struct {
		Reference *[]struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr,omitempty"`
			Type  string `xml:"type,attr,omitempty"`
			Title string `xml:"title,attr,omitempty"`
		} `xml:"reference,omitempty"`
	} `xml:"guide,omitempty"`
	Bindings *struct {
		MediaType *[]struct {
			Text      string `xml:",chardata"`
			MediaType string `xml:"media-type,attr,omitempty"`
			Handler   string `xml:"handler,attr,omitempty"`
		} `xml:"mediaType,omitempty"`
	} `xml:"bindings,omitempty"`
}

type Metadata struct {
	XMLName     xml.Name
	ID          string         `xml:"id,attr,omitempty"`
	Lang        string         `xml:"lang,attr,omitempty"`
	Dir         string         `xml:"dir,attr,omitempty"`
	Identifier  *[]Identifier  `xml:"identifier,omitempty"`
	Title       *[]Title       `xml:"title,omitempty"`
	Language    *[]Language    `xml:"language,omitempty"`
	Date        *[]Date        `xml:"date,omitempty"`
	Source      *[]Source      `xml:"source,omitempty"`
	Type        *[]Type        `xml:"type,omitempty"`
	Format      *[]Format      `xml:"format,omitempty"`
	Creator     *[]Creator     `xml:"creator,omitempty"`
	Subject     *[]Subject     `xml:"subject,omitempty"`
	Description *[]Description `xml:"description,omitempty"`
	Publisher   *[]Publisher   `xml:"publisher,omitempty"`
	Contributor *[]Contributor `xml:"contributor,omitempty"`
	Relation    *[]Relation    `xml:"relation,omitempty"`
	Coverage    *[]Coverage    `xml:"coverage,omitempty"`
	Rights      *[]Rights      `xml:"rights,omitempty"`
	Meta        *[]Meta        `xml:"meta,omitempty"`
	Link        *[]Link        `xml:"link,omitempty"`
}

type Identifier struct {
	Text   string `xml:",chardata"`
	ID     string `xml:"id,attr,omitempty"`
	Scheme string `xml:"scheme,attr,omitempty"`
}

type Title struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Language struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
}

type Date struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
}

type Source struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Type struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
}

type Format struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
}

type Creator struct {
	Text   string `xml:",chardata"`
	ID     string `xml:"id,attr,omitempty"`
	Role   string `xml:"role,attr,omitempty"`
	FileAs string `xml:"file-as,attr,omitempty"`
	Lang   string `xml:"lang,attr,omitempty"`
	Dir    string `xml:"dir,attr,omitempty"`
}

type Subject struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Description struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Publisher struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Contributor struct {
	Text   string `xml:",chardata"`
	ID     string `xml:"id,attr,omitempty"`
	Role   string `xml:"role,attr,omitempty"`
	FileAs string `xml:"file-as,attr,omitempty"`
	Lang   string `xml:"lang,attr,omitempty"`
	Dir    string `xml:"dir,attr,omitempty"`
}

type Relation struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Coverage struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Rights struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Meta struct {
	Text     string `xml:",chardata"`
	Property string `xml:"property,attr,omitempty"`
	Refines  string `xml:"refines,attr,omitempty"`
	ID       string `xml:"id,attr,omitempty"`
	Scheme   string `xml:"scheme,attr,omitempty"`
	Lang     string `xml:"lang,attr,omitempty"`
	Dir      string `xml:"dir,attr,omitempty"`
	Name     string `xml:"name,attr,omitempty"`
	Content  string `xml:"content,attr,omitempty"`
}

type Link struct {
	Text       string `xml:",chardata"`
	Href       string `xml:"href,attr,omitempty"`
	Rel        string `xml:"rel,attr,omitempty"`
	ID         string `xml:"id,attr,omitempty"`
	Refines    string `xml:"refines,attr,omitempty"`
	MediaType  string `xml:"media-type,attr,omitempty"`
	Properties string `xml:"properties,attr,omitempty"`
}
