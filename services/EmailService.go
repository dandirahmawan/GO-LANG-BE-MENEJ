package services

import (
	"github.com/dandirahmawan/menej_api_go/config"
	"github.com/dandirahmawan/menej_api_go/model"
)

func SendEmailInvitation(data model.ViewInvitationModel) {
	var to []string
	var cc []string

	/*set up recipient*/
	email := data.InvitationEmail
	to = append(to, email)

	msg := MessageEmailInvitation(data.UserNameInviting, data.InvitationId)
	config.SendMail(to, cc, "Invitation", msg)
}

func MessageEmailInvitation(userName, codeConfirm string) string {
	var message = "<div style='background: #e1e1e1; padding: 40px'>\n" +
		"                <div style='max-width: 280px; padding: 10px; margin: auto; border: 1px solid #d8d8d8; border-radius: 5px; background: #FFF;'>\n" +
		"                    <table>\n" +
		"                        <tr> \n" +
		"                            <td><img src='https://scontent.fcgk4-4.fna.fbcdn.net/v/t1.0-9/119041998_3359216697526275_5140025698278560974_n.jpg?_nc_cat=101&_nc_sid=730e14&_nc_eui2=AeGb5dgJggtCxe8Gf7JhczaLZ-yXZaga0V5n7JdlqBrRXr4_7jq93tmJrFeJN79mAbmFC1qXoeXqwvFKSueo-xBW&_nc_ohc=I5ndHSrf8IYAX_mEL6u&_nc_ht=scontent.fcgk4-4.fna&oh=74b8358dc5aa99eb7df0b4f901459243&oe=5F7A057C'/></td>\n" +
		"                        </tr>\n" +
		"                        <tr> \n" +
		"                            <td>\n" +
		"                                <div style='font-size: 20px'><strong>Invitation</strong></div>\n" +
		"                            </td>\n" +
		"                        </tr>\n" +
		"                        <tr>\n" +
		"                            <td>\n" +
		"                                <strong style='font-size: 14px'>Hi, i am " + userName + "</strong><br/>\n" +
		"                                <div style='font-size: 11px'>\n" +
		"                                    I want invite you to colaborating in menej for project <strong>Menej development</strong><br/>\n" +
		"                                </div>\n" +
		"                            </td>\n" +
		"                        </tr>\n" +
		"                        <tr>\n" +
		"                            <td> \n" +
		"                                <a href='http://localhost:3006/invitation?conf=" + codeConfirm + "'>\n" +
		"                                    <button style='padding: 7px; border: none; background: #386384; color: #FFF; border-radius: 3px; font-size: 11px;margin-top: 10px'>\n" +
		"                                        Accept\n" +
		"                                    </button>\n" +
		"                                </a>\n" +
		"                            </td>\n" +
		"                        </tr> \n" +
		"                        <tr>\n" +
		"                            <td> \n" +
		"								 <div style='border-top: 1px solid #CCC; color: #a0a0a0; padding-top: 10px; font-size: 11px; margin-top: 10px'>" +
		"                                	<div>\n" +
		"                                    This email will expired in 24 hours, click accept for the next step of approvement" +
		"								 	</div>\n" +
		"								 </div>\n" +
		"                            </td>\n" +
		"                        </tr>\n" +
		"                    </table>\n" +
		"                </div>\n" +
		"            </div>"

	return message
}
