package main

import (
	"fmt"
	"github.com/BrianLeishman/go-imap"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//imap.Verbose = true
	imap.RetryCount = 0
	im, err := imap.New("support@klf.ch", "9*T8ckr1", "admin.yop.ch", 993)
	check(err)
	defer im.Close()

	folders, err := im.GetFolders()
	check(err)

	for _, f := range folders {
		fmt.Println("Working on " + f)
		err = im.SelectFolder(f)
		check(err)

		uids, err := im.GetUIDs("ALL")
		check(err)

		emails, err := im.GetEmails(uids...)
		check(err)

		if "INBOX" == f && len(emails) != 0 {
			fmt.Print(emails)
		}
	}

}

// SelectFolder selects a folder
func (d *imap.Dialer) SelectFolder(folder string) (err error) {
	_, err = d.Exec(`EXAMINE "`+AddSlashes.Replace(folder)+`"`, true, RetryCount, nil)
	if err != nil {
		return
	}
	d.Folder = folder
	return nil
}
func MarkMessageAsRead(d *imap.Dialer) {

}
