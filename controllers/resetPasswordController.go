package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"main.go/database"
	"main.go/models"
	"math/rand"
	"net/smtp"
	"os"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	if err := database.DB.Find(&user, "email = ?", data["email"]).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": "There is an error in finding email method",
		})
	}

	if user.Email == data["email"] {

		email := data["email"]
		token := RandStingRunes(100)

		passwordReset := models.PasswordReset{
			Email: email,
			Token: token,
		}

		database.DB.Create(&passwordReset)

		from := "help@minibell.com"

		to := []string{
			email,
		}

		url := "http://localhost:3000/signin/resetpass/" + token

		message := []byte("Click <a href=\"" + url + "\">here</a> to reset your password")

		err := smtp.SendMail("0.0.0.0:1025", nil, from, to, message)

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "check your emails",
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

}

func Reset(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return c.Status(400).JSON(fiber.Map{
			"message": "password does not match",
		})
	}

	passwordReset := models.PasswordReset{}

	database.DB.Where("token = ?", data["token"]).Last(&passwordReset)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func VerifyMail(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	email := data["email"]
	name := data["firstName"]

	subject := "Subject: Confirmation\n"

	godotenv.Load()
	from := os.Getenv("EMAIL_ADDRESS")

	to := []string{
		email,
	}

	url := "http://localhost:3000/signin/"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "" +
		"<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\"><html data-editor-version=\"2\" class=\"sg-campaigns\" xmlns=\"http://www.w3.org/1999/xhtml\"><head>\n      " +
		"<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\n      " +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1\">\n      " +
		"<!--[if !mso]><!-->\n      " +
		"<meta http-equiv=\"X-UA-Compatible\" content=\"IE=Edge\">\n      " +
		"<!--<![endif]-->\n      " +
		"<!--[if (gte mso 9)|(IE)]>\n      " +
		"<xml>\n        " +
		"<o:OfficeDocumentSettings>\n          " +
		"<o:AllowPNG/>\n          " +
		"<o:PixelsPerInch>96</o:PixelsPerInch>\n        " +
		"</o:OfficeDocumentSettings>\n      " +
		"</xml>\n      " +
		"<![endif]-->\n      " +
		"<!--[if (gte mso 9)|(IE)]>\n  " +
		"<style type=\"text/css\">\n    " +
		"body {width: 600px;margin: 0 auto;}\n    " +
		"table {border-collapse: collapse;}\n    " +
		"table, td {mso-table-lspace: 0pt;mso-table-rspace: 0pt;}\n    " +
		"img {-ms-interpolation-mode: bicubic;}\n  " +
		"</style>\n<![endif]-->\n      " +
		"<style type=\"text/css\">\n    " +
		"body, p, div {\n      " +
		"font-family: inherit;\n      " +
		"font-size: 14px;\n    }\n    " +
		"body {\n      " +
		"color: #000000;\n    }\n    " +
		"body a {\n      " +
		"color: #1188E6;\n      " +
		"text-decoration: none;\n    }\n    " +
		"p { margin: 0; padding: 0; }\n    " +
		"table.wrapper {\n      " +
		"width:100% !important;\n      " +
		"table-layout: fixed;\n      " +
		"-webkit-font-smoothing: antialiased;\n      " +
		"-webkit-text-size-adjust: 100%;\n      " +
		"-moz-text-size-adjust: 100%;\n      " +
		"-ms-text-size-adjust: 100%;\n    }\n    " +
		"img.max-width {\n      " +
		"max-width: 100% !important;\n    }\n    " +
		".column.of-2 {\n      " +
		"width: 50%;\n    }\n    " +
		".column.of-3 {\n      " +
		"width: 33.333%;\n    }\n    " +
		".column.of-4 {\n      width: 25%;\n    }\n    " +
		"@media screen and (max-width:480px) {\n      " +
		".preheader .rightColumnContent,\n      " +
		".footer .rightColumnContent {\n        " +
		"text-align: left !important;\n      }\n      " +
		".preheader .rightColumnContent div,\n      ." +
		"preheader .rightColumnContent span,\n      " +
		".footer .rightColumnContent div,\n      " +
		".footer .rightColumnContent span {\n        " +
		"text-align: left !important;\n      }\n      " +
		".preheader .rightColumnContent,\n      " +
		".preheader .leftColumnContent {\n        " +
		"font-size: 80% !important;\n        " +
		"padding: 5px 0;\n      }\n      " +
		"table.wrapper-mobile {\n        " +
		"width: 100% !important;\n       " +
		" table-layout: fixed;\n      }\n      " +
		"img.max-width {\n        " +
		"height: auto !important;\n        " +
		"max-width: 100% !important;\n      }\n     " +
		" a.bulletproof-button {\n        " +
		"display: block !important;\n       " +
		" width: auto !important;\n        " +
		"font-size: 80%;\n       " +
		" padding-left: 0 !important;\n        " +
		"padding-right: 0 !important;\n      }\n      " +
		".columns {\n        width: 100% !important;\n      }\n      " +
		".column {\n        display: block !important;\n        " +
		"width: 100% !important;\n        padding-left: 0 !important;\n        " +
		"padding-right: 0 !important;\n        margin-left: 0 !important;\n        " +
		"margin-right: 0 !important;\n      }\n    }\n  </style>\n      " +
		"<!--user entered Head Start-->" +
		"<link href=\"https://fonts.googleapis.com/css?family=Muli&display=swap\" " +
		"rel=\"stylesheet\">" +
		"<style>\nbody {font-family: 'Muli', sans-serif;}\n</style>" +
		"<!--End Head user entered-->\n    </head>\n    " +
		"<body>\n      " +
		"<center class=\"wrapper\" data-link-color=\"#1188E6\" data-body-style=\"font-size:14px; font-family:inherit; color:#000000; background-color:#FFFFFF;\">\n        " +
		"<div class=\"webkit\">\n          " +
		"<table cellpadding=\"0\" cellspacing=\"0\" border=\"0\" width=\"100%\" class=\"wrapper\" bgcolor=\"#FFFFFF\">\n            " +
		"<tbody><tr>\n              " +
		"<td valign=\"top\" bgcolor=\"#FFFFFF\" width=\"100%\">\n                " +
		"<table width=\"100%\" role=\"content-container\" class=\"outer\" align=\"center\" cellpadding=\"0\" cellspacing=\"0\" border=\"0\">\n                  " +
		"<tbody><tr>\n                    <td width=\"100%\">\n                      " +
		"<table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" border=\"0\">\n                        <tbody><tr>\n                          " +
		"<td>\n                            <!--[if mso]>\n    <center>\n    <table><tr><td width=\"600\">\n  <![endif]-->\n                                    " +
		"<table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" border=\"0\" style=\"width:100%; max-width:600px;\" align=\"center\">\n                                     " +
		" <tbody><tr>\n                                        " +
		"<td role=\"modules-container\" style=\"padding:0px 0px 0px 0px; color:#000000; text-align:left;\" bgcolor=\"#FFFFFF\" width=\"100%\" align=\"left\"><table class=\"module preheader preheader-hide\" role=\"module\" data-type=\"preheader\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"display: none !important; mso-hide: all; visibility: hidden; opacity: 0; color: transparent; height: 0; width: 0;\">\n    " +
		"<tbody><tr>\n      <td role=\"module-content\">\n        <p></p>\n      </td>\n    </tr>\n  " +
		"</tbody></table><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" align=\"center\" width=\"100%\" role=\"module\" data-type=\"columns\" style=\"padding:30px 20px 30px 20px;\" bgcolor=\"#f6f6f6\">\n    " +
		"<tbody>\n      <tr role=\"module-content\">\n        <td height=\"100%\" valign=\"top\">\n          " +
		"<table class=\"column\" width=\"540\" style=\"width:540px; border-spacing:0; border-collapse:collapse; margin:0px 10px 0px 10px;\" cellpadding=\"0\" cellspacing=\"0\" align=\"left\" border=\"0\" bgcolor=\"\">\n            " +
		"<tbody>\n              <tr>\n                " +
		"<td style=\"padding:0px;margin:0px;border-spacing:0;\"><table class=\"wrapper\" role=\"module\" data-type=\"image\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"72aac1ba-9036-4a77-b9d5-9a60d9b05cba\">\n    " +
		"<tbody>\n      <tr>\n        <td style=\"font-size:6px; line-height:10px; padding:0px 0px 0px 0px;\" valign=\"top\" align=\"center\">\n          " +
		"<img class=\"max-width\" border=\"0\" style=\"display:block; color:#000000; text-decoration:none; font-family:Helvetica, arial, sans-serif; font-size:16px;\" width=\"29\" alt=\"\" data-proportionally-constrained=\"true\" data-responsive=\"false\" src=\"http://cdn.mcauto-images-production.sendgrid.net/954c252fedab403f/9200c1c9-b1bd-47ed-993c-ee2950a0f239/29x27.png\" height=\"27\">\n        " +
		"</td>\n      </tr>\n    </tbody>\n  " +
		"</table><table class=\"module\" role=\"module\" data-type=\"spacer\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"331cde94-eb45-45dc-8852-b7dbeb9101d7\">\n    <tbody>\n      " +
		"<tr>\n        <td style=\"padding:0px 0px 20px 0px;\" role=\"module-content\" bgcolor=\"\">\n        </td>\n      </tr>\n    </tbody>\n  " +
		"</table><table class=\"wrapper\" role=\"module\" data-type=\"image\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"d8508015-a2cb-488c-9877-d46adf313282\">\n    <tbody>\n      <tr>\n        " +
		"<td style=\"font-size:6px; line-height:10px; padding:0px 0px 0px 0px;\" valign=\"top\" align=\"center\">\n          " +
		"<img class=\"max-width\" border=\"0\" style=\"display:block; color:#000000; text-decoration:none; font-family:Helvetica, arial, sans-serif; font-size:16px;\" width=\"95\" alt=\"\" data-proportionally-constrained=\"true\" data-responsive=\"false\" src=\"http://cdn.mcauto-images-production.sendgrid.net/954c252fedab403f/61156dfa-7b7f-4020-85f8-a586addf4288/95x33.png\" height=\"33\">\n        " +
		"</td>\n      </tr>\n    </tbody>\n  " +
		"</table><table class=\"module\" role=\"module\" data-type=\"spacer\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"27716fe9-ee64-4a64-94f9-a4f28bc172a0\">\n    " +
		"<tbody>\n      <tr>\n        <td style=\"padding:0px 0px 30px 0px;\" role=\"module-content\" bgcolor=\"\">\n        </td>\n      </tr>\n    </tbody>\n  " +
		"</table><table class=\"module\" role=\"module\" data-type=\"text\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"948e3f3f-5214-4721-a90e-625a47b1c957\" data-mc-module-version=\"2019-10-22\">\n    " +
		"<tbody>\n      <tr>\n        " +
		"<td style=\"padding:50px 30px 18px 30px; line-height:36px; text-align:inherit; background-color:#ffffff;\" height=\"100%\" valign=\"top\" bgcolor=\"#ffffff\" role=\"module-content\"><div><div style=\"font-family: inherit; text-align: center\"><span style=\"font-size: 43px\">Thanks for signing up, " +
		name +
		"!&nbsp;</span></div><div></div></div></td>\n      " +
		"</tr>\n    </tbody>\n  " +
		"</table><table class=\"module\" role=\"module\" data-type=\"text\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"a10dcb57-ad22-4f4d-b765-1d427dfddb4e\" data-mc-module-version=\"2019-10-22\">\n    " +
		"<tbody>\n      <tr>\n        " +
		"<td style=\"padding:18px 30px 18px 30px; line-height:22px; text-align:inherit; background-color:#ffffff;\" height=\"100%\" valign=\"top\" bgcolor=\"#ffffff\" role=\"module-content\"><div><div style=\"font-family: inherit; text-align: center\"><span style=\"font-size: 18px\">Please verify your email address to" +
		"</span>" +
		"<span style=\"color: #000000; font-size: 18px; font-family: arial,helvetica,sans-serif\">" +
		" get access to thousands of exclusive items" +
		"</span><span style=\"font-size: 18px\">." +
		"</span></div>\n<div style=\"font-family: inherit; text-align: center\"><span style=\"color: #ffbe00; font-size: 18px\"><strong>Thank you!&nbsp;" +
		"</strong></span></div><div></div></div></td>\n      </tr>\n    " +
		"</tbody>\n  " +
		"</table><table class=\"module\" role=\"module\" data-type=\"spacer\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"7770fdab-634a-4f62-a277-1c66b2646d8d\">\n    " +
		"<tbody>\n      <tr>\n        <td style=\"padding:0px 0px 20px 0px;\" role=\"module-content\" bgcolor=\"#ffffff\">\n        </td>\n      </tr>\n    " +
		"</tbody>\n  " +
		"</table><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"module\" data-role=\"module-button\" data-type=\"button\" role=\"module\" style=\"table-layout:fixed;\" width=\"100%\" data-muid=\"d050540f-4672-4f31-80d9-b395dc08abe1\">\n      " +
		"<tbody>\n        <tr>\n          <td align=\"center\" bgcolor=\"#ffffff\" class=\"outer-td\" style=\"padding:0px 0px 0px 0px;\">\n            " +
		"<table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"wrapper-mobile\" style=\"text-align:center;\">\n              <tbody>\n                " +
		"<tr>\n                " +
		"<td align=\"center\" bgcolor=\"#ffbe00\" class=\"inner-td\" style=\"border-radius:6px; font-size:16px; text-align:center; background-color:inherit;\">\n                  " +
		"<a href=\"" +
		url +
		"\" style=\"background-color:#ffbe00; border:1px solid #ffbe00; border-color:#ffbe00; border-radius:0px; border-width:1px; color:#000000; display:inline-block; font-size:14px; font-weight:normal; letter-spacing:0px; line-height:normal; padding:12px 40px 12px 40px; text-align:center; text-decoration:none; border-style:solid; font-family:inherit;\" target=\"_blank\">" +
		"Verify Email Now</a>\n                </td>\n                </tr>\n              </tbody>\n            " +
		"</table>\n          </td>\n        </tr>\n      </tbody>\n    " +
		"</table><table class=\"module\" role=\"module\" data-type=\"spacer\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"7770fdab-634a-4f62-a277-1c66b2646d8d.1\">\n    " +
		"<tbody>\n      <tr>\n        <td style=\"padding:0px 0px 50px 0px;\" role=\"module-content\" bgcolor=\"#ffffff\">\n        </td>\n      </tr>\n    " +
		"</tbody>\n  </table>" +
		"<table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"module\" data-role=\"module-button\" data-type=\"button\" role=\"module\" style=\"table-layout:fixed;\" width=\"100%\" data-muid=\"d050540f-4672-4f31-80d9-b395dc08abe1.1\">\n      <tbody>\n        <tr>\n          <td align=\"center\" bgcolor=\"#6e6e6e\" class=\"outer-td\" style=\"padding:0px 0px 0px 0px;\">\n            <table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"wrapper-mobile\" style=\"text-align:center;\">\n              <tbody>\n                <tr>\n                " +
		"<td align=\"center\" bgcolor=\"#ffbe00\" class=\"inner-td\" style=\"border-radius:6px; font-size:16px; text-align:center; background-color:inherit;\">\n    </br>              " +
		"<a href=\"\" style=\"background-color:#ffbe00; border:1px solid #ffbe00; border-color:#ffbe00; border-radius:0px; border-width:1px; color:#000000; display:inline-block; font-size:14px; font-weight:normal; letter-spacing:0px; line-height:normal; padding:12px 40px 12px 40px; text-align:center; text-decoration:none; border-style:solid; font-family:inherit;\" target=\"_blank\">Contact Support</a>\n                " +
		"</td>\n                </tr>\n              </tbody>\n            </table>\n          </td>\n        </tr>\n      </tbody>\n    " +
		"</table><table class=\"module\" role=\"module\" data-type=\"spacer\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"table-layout: fixed;\" data-muid=\"c37cc5b7-79f4-4ac8-b825-9645974c984e\">\n    " +
		"<tbody>\n      <tr>\n        " +
		"<td style=\"padding:0px 0px 30px 0px;\" role=\"module-content\" bgcolor=\"6E6E6E\">\n        " +
		"</td>\n      </tr>\n    </tbody>\n  </table></td>\n              </tr>\n            </tbody>\n          " +
		"</table>\n          \n        </td>\n      </tr>\n    </tbody>\n  " +
		"</table><div data-role=\"module-unsubscribe\" class=\"module\" role=\"module\" data-type=\"unsubscribe\" style=\"color:#444444; font-size:12px; line-height:20px; padding:16px 16px 16px 16px; text-align:Center;\" data-muid=\"4e838cf3-9892-4a6d-94d6-170e474d21e5\">\n                                            " +
		"<div class=\"Unsubscribe--addressLine\"><p class=\"Unsubscribe--senderName\" style=\"font-size:12px; line-height:20px;\">" +
		"mini Bell" +
		"</p><p style=\"font-size:12px; line-height:20px;\"><span class=\"Unsubscribe--senderAddress\">" +
		"No: 15, Mount Lavinia" +
		"</span>, <span class=\"Unsubscribe--senderCity\">" +
		"Colombo" +
		"</span>, <span class=\"Unsubscribe--senderState\">" +
		"Sri lanka" +
		"</span> </p></div>\n                                            " +
		"<p style=\"font-size:12px; line-height:20px;\"><a class=\"Unsubscribe--unsubscribeLink\" href=\"{{{unsubscribe}}}\" target=\"_blank\" style=\"\">Unsubscribe</a> - " +
		"<a href=\"{{{unsubscribe_preferences}}}\" target=\"_blank\" class=\"Unsubscribe--unsubscribePreferences\" style=\"\">Unsubscribe Preferences</a></p>\n                                          </div><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"module\" data-role=\"module-button\" data-type=\"button\" role=\"module\" style=\"table-layout:fixed;\" width=\"100%\" data-muid=\"550f60a9-c478-496c-b705-077cf7b1ba9a\">\n      <tbody>\n        <tr>\n          <td align=\"center\" bgcolor=\"\" class=\"outer-td\" style=\"padding:0px 0px 20px 0px;\">\n            <table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"wrapper-mobile\" style=\"text-align:center;\">\n              <tbody>\n                <tr>\n                <td align=\"center\" bgcolor=\"#f5f8fd\" class=\"inner-td\" style=\"border-radius:6px; font-size:16px; text-align:center; background-color:inherit;\"><a href=\"https://sendgrid.com/\" style=\"background-color:#f5f8fd; border:1px solid #f5f8fd; border-color:#f5f8fd; border-radius:25px; border-width:1px; color:#a8b9d5; display:inline-block; font-size:10px; font-weight:normal; letter-spacing:0px; line-height:normal; padding:5px 18px 5px 18px; text-align:center; text-decoration:none; border-style:solid; font-family:helvetica,sans-serif;\" target=\"_blank\">♥ POWERED BY TWILIO SENDGRID</a></td>\n                </tr>\n              </tbody>\n            </table>\n          </td>\n        </tr>\n      </tbody>\n    </table></td>\n                                      </tr>\n                                    </tbody></table>\n                                    <!--[if mso]>\n                                  </td>\n                                </tr>\n                              </table>\n                            </center>\n                            <![endif]-->\n                          </td>\n                        </tr>\n                      </tbody></table>\n                    </td>\n                  </tr>\n                </tbody></table>\n              </td>\n            </tr>\n          </tbody></table>\n        </div>\n      </center>\n    \n  \n</body></html>"

	msg := []byte(subject + mime + body)

	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, msg)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "check your email",
	})
}

func RandStingRunes(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func SendEmail(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "body parsing",
		})
	}

	email := data["emailaddress"]
	name := data["companyname"]

	subject := "Subject: Confirmation\n"

	godotenv.Load()
	from := os.Getenv("EMAIL_ADDRESS")

	to := []string{
		email,
	}

	url := "http://localhost:3000/signin/"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body := "<!DOCTYPE html>\n" +
		"<html lang=\"en\">\n  " +
		"<head>\n    " +
		"<meta charset=\"UTF-8\" />\n    " +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n    " +
		"<title>Document</title>\n    " +
		"<style>\n      " +
		"@import url(\"https://fonts.googleapis.com/css2?family=Raleway:ital,wght@1,200&display=swap\");\n\n      " +
		"* {\n        margin: 0;\n        padding: 0;\n        border: 0;\n      }\n\n      " +
		"body {\n        font-family: 'Montserrat';\n        background-color: whitesmoke;\n        font-size: 19px;\n        max-width: 800px;\n        margin: 0 auto;\n        padding: 3%;\n      }\n\n      " +
		"img {\n        max-width: 100%;\n      }\n\n      " +
		"header {\n        width: 98%;\n      }\n\n      " +
		"#logo {\n        max-width: 120px;\n        margin: 3% 0 3% 3%;\n        float: left;\n      }\n\n      " +
		"#wrapper {\n        background-color: #f0f6fb;\n      }\n\n      " +
		"#social {\n        float: right;\n        margin: 3% 2% 4% 3%;\n        list-style-type: none;\n      }\n\n      " +
		"#social > li {\n        display: inline;\n      }\n\n      " +
		"#social > li > a > img {\n        max-width: 35px;\n      }\n\n      " +
		"h1,\n      p {\n        text-align: center;\n        margin: 3%;\n      }\n      " +
		".btn {\n        float: right;\n        margin: 0 2% 4% 0;\n        background-color: #B5D6DE;\n        color: #f6faff;\n        text-decoration: none;\n        font-weight: 800;\n        padding: 8px 12px;\n        border-radius: 8px;\n        letter-spacing: 2px;\n      }\n\n      " +
		"hr {\n        height: 1px;\n        background-color: #303840;\n        clear: both;\n        width: 96%;\n        margin: auto;\n      }\n\n      " +
		"#contact {\n        text-align: center;\n        padding-bottom: 3%;\n        line-height: 16px;\n        font-size: 12px;\n        color: #303840;\n      }\n    " +
		"</style>\n  " +
		"</head>\n  " +
		"<body>\n    " +
		"<div id=\"wrapper\">\n      " +
		"<header>\n        " +
		"<div id=\"logo\">\n          " +
		"<img\n            src=\"logo.png\"\n            alt=\"\"\n          />\n        " +
		"</div>\n        " +
		"<div>\n          " +
		"<ul id=\"social\">\n            " +
		"<li>\n              " +
		"<a href=\"#\" target=\"_blank\"\n                >" +
		"<img\n                  src=\"facebook.png\"\n                  alt=\"\"\n              />" +
		"</a>\n            " +
		"</li>\n            " +
		"<li>\n              " +
		"<a href=\"#\" target=\"_blank\"\n                >" +
		"<img\n                  src=\"instagram.png\"\n                  alt=\"\"\n              />" +
		"</a>\n            " +
		"</li>\n            " +
		"<li>\n              " +
		"<a href=\"#\" target=\"_blank\"\n                >" +
		"<img\n                  src=\"twitter.png\"\n                  alt=\"\"\n              />" +
		"</a>\n            " +
		"</li>\n          " +
		"</ul>\n        " +
		"</div>\n      " +
		"</header>\n      " +
		"<div id=\"banner\">\n        " +
		"<img\n          src=\"https://thumbs.dreamstime.com/z/rainbow-love-heart-background-red-wood-60045149.jpg\"\n          alt=\"\"\n        />\n      " +
		"</div>\n      " +
		"<div class=\"one-col\">\n        " +
		"<h1>" +
		"Congratulations on Creating Your Company!" +
		"</h1>\n\n        " +
		"<p>\n          Your Company  <b>" +
		name +
		"</b> has been added to MiniBell.\n        </p>\n\n        " +
		"<p>\n          Now you can add products and let our cutomers explore them and buy them.\n        " +
		"</p>\n\n        " +
		"<p>" +
		"<b>\n            Happy Selling!\n          </b>" +
		"</p>\n\n        <a href=" +
		url +
		"" +
		" class=\"btn\">Add Products</a>\n\n        " +
		"<hr />\n\n        " +
		"<footer>\n          " +
		"<p id=\"contact\">\n            Stay connected <br />\n            help@minibell.com<br />\n            +94 112 783 789 <br />\n          " +
		"</p>\n        " +
		"</footer>\n      " +
		"</div>\n    " +
		"</div>\n  " +
		"</body>\n" +
		"</html>"

	msg := []byte(subject + mime + body)

	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, msg)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "keyFunc",
		})
	}

	return c.JSON(fiber.Map{
		"message": "check your email",
	})
}
