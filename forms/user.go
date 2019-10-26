package forms

// CreateUserCommand Struct
type CreateUserCommand struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

// LoginCommand Struct
type LoginCommand struct {
	Email    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ForgetAPI func
type ForgetAPI struct {
	Email string `json:"username" binding:"required"`
}

// UpdateUserCommand Struct
type UpdateUserCommand struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

//DeleteConf struct
type DeleteConf struct {
	Bid string `json:"bid" binding:"required"`
}

// GroupConf struct
type GroupConf struct {
	Gid                 string `json:"gid" binding:"required"`
	Bid                 string `json:"bid" binding:"required"`
	Eid                 string `json:"eid" binding:"required"`
	Groupname           string `json:"groupname" binding:"required"`
	Prinumber           string `json:"prinumber" binding:"required"`
	Rules               string `json:"rules" binding:"required"`
	Keyword             string `json:"keyword" binding:"required"`
	Timeout             string `json:"timeout" binding:"required"`
	Hdaytext            string `json:"hdaytext" binding:"required"`
	Hdayaudio           string `json:"hdayaudio" binding:"required"`
	Replymessage        string `json:"replymessage" binding:"required"`
	LeadAction          string `json:"leadaction" binding:"required"`
	Voicemessage        string `json:"voicemessage" binding:"required"`
	Voicemessagetext    string `json:"voicemessagetext" binding:"required"`
	Greetings           string `json:"greetings" binding:"required"`
	URL                 string `json:"url" binding:"required"`
	Bday                string `json:"bday" binding:"required"`
	AccessReport        string `json:"access_reports" binding:"required"`
	OnCallAction        string `json:"oncallaction" binding:"required"`
	LandingRegion       string `json:"landingregion" binding:"required"`
	OnEditAction        string `json:"oneditaction" binding:"required"`
	OnHangup            string `json:"onhangup" binding:"required"`
	ConnectOwner        string `json:"connectowner" binding:"required"`
	Replyattmsg         string `json:"replyattmsg" binding:"required"`
	Pincode             string `json:"pincode" binding:"required"`
	SupportAction       string `json:"supportaction" binding:"required"`
	ReplyToCustomer     string `json:"replytocustomer" binding:"required"`
	ReplyToCustMissed   string `json:"replytocust_missed" binding:"required"`
	ReplyToCustRepCal   string `json:"replytocust_repcal" binding:"required"`
	ReplyToCustAttended string `json:"replytocust_attended" binding:"required"`
	RplyToCustNonBusCal string `json:"replytocust_nonbuscall" binding:"required"`
	Replytocustvoice    string `json:"replytocust_voice" binding:"required"`
	ReplyToExecutive    string `json:"replytoexecutive" binding:"required"`
	ReplyExtAttended    string `json:"replyext_attended" binding:"required"`
	ReplyExtMissed      string `json:"replyext_missed" binding:"required"`
	RecordNotice        string `json:"recordnotice" binding:"required"`
	Record              string `json:"record" binding:"required"`
	MailAlert           string `json:"mailalert" binding:"required"`
	ReplyToCustVoiText  string `json:"replytocust_voitext" binding:"required"`
	SameExe             string `json:"sameexe" binding:"required"`
	AllGroup            string `json:"allgroup" binding:"required"`
	MissCall            string `json:"misscall" binding:"required"`
	Mailalerttoowner    string `json:"mailalerttowoner" binding:"required"`
	CustomVoiceFile     string `json:"custom_voice_file" binding:"required"`
	PrimaryRule         string `json:"primary_rule" binding:"required"`
	SupportGrp          string `json:"supportgrp" binding:"required"`
	Status              string `json:"status" binding:"required"`
}
