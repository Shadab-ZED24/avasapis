package model

import (
	"encoding/base64"

	"avasapis/db"
	"avasapis/forms"

	"gopkg.in/mgo.v2/bson"
)

// type User struct {
// 	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
// 	Firstname string        `json:"firstname"`
// 	Lastname  string        `json:"lastname"`
// }

// User Struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserModel (Model) Struct
type UserModel struct{}

var (
	dbConnect      = db.NewConnection("localhost")
	dbName         = "avas"
	collectionName = "user"
)

//Login user
func (m *UserModel) Login(email string) (user User, err error) {
	collection := dbConnect.Use(dbName, collectionName)
	query := []bson.M{{
		"$lookup": bson.M{ // lookup the documents table here
			"from":         "business",
			"localField":   "bid",
			"foreignField": "bid",
			"as":           "businessDetails",
		}},
		{"$unwind": "$businessDetails"},
		{"$match": bson.M{"username": email}}}

	pipe := collection.Pipe(query)
	err = pipe.One(&user)
	return
}

//Forget func
// func (m *UserModel) Forget(email string, pass string) (err error) {
// 	collection := dbConnect.Use(dbName, collectionName)
// 	query := bson.M{
// 				"username": email,
// 			}, bson.M{"$set": bson.M{
// 				"password": pass,
// 			}}
// 	collection.Update(query)
// 	return
// }

//Add Group Config
func (m *UserModel) Add(data forms.GroupConf) error {
	collection := dbConnect.Use(dbName, data.Bid+"_groups")
	groupkey := base64.StdEncoding.EncodeToString([]byte(data.Bid + "_" + data.Gid))
	err := collection.Insert(bson.M{
		"_id":                    data.Gid,
		"gid":                    data.Gid,
		"bid":                    data.Bid,
		"eid":                    data.Eid,
		"groupname":              data.Groupname,
		"prinumber":              data.Prinumber,
		"rules":                  data.Rules,
		"keyword":                data.Keyword,
		"timeout":                data.Timeout,
		"hdaytext":               data.Hdaytext,
		"hdayaudio":              data.Hdayaudio,
		"replymessage":           data.Replymessage,
		"leadaction":             data.LeadAction,
		"voicemessage":           data.Voicemessage,
		"voicemessagetext":       data.Voicemessagetext,
		"greetings":              data.Greetings,
		"url":                    data.URL,
		"bday":                   data.Bday,
		"access_reports":         data.AccessReport,
		"oncallaction":           data.OnCallAction,
		"landingregion":          data.LandingRegion,
		"oneditaction":           data.OnEditAction,
		"onhangup":               data.OnHangup,
		"connectowner":           data.ConnectOwner,
		"replyattmsg":            data.Replyattmsg,
		"pincode":                data.Pincode,
		"supportaction":          data.SupportAction,
		"replytocustomer":        data.ReplyToCustomer,
		"replytocust_missed":     data.ReplyToCustMissed,
		"replytocust_repcal":     data.ReplyToCustRepCal,
		"replytocust_attended":   data.ReplyToCustAttended,
		"replytocust_nonbuscall": data.RplyToCustNonBusCal,
		"replytoexecutive":       data.ReplyToExecutive,
		"replytocust_voice":      data.Replytocustvoice,
		"replyext_attended":      data.ReplyExtAttended,
		"replyext_missed":        data.ReplyExtMissed,
		"recordnotice":           data.RecordNotice,
		"record":                 data.Record,
		"mailalert":              data.MailAlert,
		"replytocust_voitext":    data.ReplyToCustVoiText,
		"sameexe":                data.SameExe,
		"allgroup":               data.AllGroup,
		"misscall":               data.MissCall,
		"mailalerttowoner":       data.Mailalerttoowner,
		"custom_voice_file":      data.CustomVoiceFile,
		"groupkey":               groupkey,
		"primary_rule":           data.PrimaryRule,
		"supportgrp":             data.SupportGrp,
		"status":                 data.Status,
	})
	return err
}

//UpdateConf user
func (m *UserModel) UpdateConf(id string, data forms.GroupConf) (err error) {
	collection := dbConnect.Use(dbName, data.Bid+"_groups")
	err = collection.UpdateId(id, data)

	return
}

//DeleteConf user
func (m *UserModel) DeleteConf(id string, data forms.DeleteConf) (err error) {
	collection := dbConnect.Use(dbName, data.Bid+"_groups")
	query := bson.M{"$set": bson.M{"status": "0"}}
	err = collection.UpdateId(id, query)

	return
}

// Create a user
func (m *UserModel) Create(data forms.CreateUserCommand) error {
	collection := dbConnect.Use(dbName, collectionName)
	err := collection.Insert(bson.M{
		"firstname": data.Firstname,
		"lastname":  data.Lastname,
	})
	return err
}

// GetAll users
func (m *UserModel) GetAll() (list []User, err error) {
	collection := dbConnect.Use(dbName, collectionName)
	err = collection.Find(bson.M{}).All(&list)
	return
}

//GetOne user
func (m *UserModel) GetOne(id string) (user User, err error) {
	collection := dbConnect.Use(dbName, collectionName)
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return
}

//UpdateOne user
func (m *UserModel) UpdateOne(id string, data forms.UpdateUserCommand) (err error) {
	collection := dbConnect.Use(dbName, collectionName)
	err = collection.UpdateId(bson.ObjectIdHex(id), data)

	return
}

//DeleteOne user
func (m *UserModel) DeleteOne(id string) (err error) {
	collection := dbConnect.Use(dbName, collectionName)
	err = collection.RemoveId(bson.ObjectIdHex(id))

	return
}
