package internal

import "fmt"

// あなたは簡単な住所帳アプリケーションを作成しようとしています。このアプリケーションでは、以下の機能を持つ住所帳を管理する必要があります。

// 1. 住所帳に新しい連絡先を追加する
// 2. 連絡先をIDで検索する
// 3. すべての連絡先を表示する

type Contact struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

// map[int]はContactを扱う際のキーとしてintを使うから
// これのおかげでaddress.Contacts[contact.ID]とかが使える
// キーなしのスライスでも行けなくはないが、こっちの方が簡単
// intじゃなくてstringとかでもいけるけど、管理が面倒だから大体intかも
type AddressBook struct {
	Contacts map[int]Contact
}

func NewAddressBook() *AddressBook {
	return &AddressBook{
		Contacts: make(map[int]Contact),
	}
}

func (address *AddressBook) AddContact(contact Contact) {
	address.Contacts[contact.ID] = contact
}

func (address *AddressBook) GetContact(id int) (Contact, bool) {
	contact, exists := address.Contacts[id]
	return contact, exists
}

func (address *AddressBook) ListContact() []Contact {
	contactList := []Contact{}
	for _, contact := range address.Contacts {
		contactList = append(contactList, contact)
	}
	return contactList
}

func StructTest() {
	address := NewAddressBook()
	address.AddContact(Contact{ID: 1, Name: "hoge", Address: "Tokyo", Phone: "0120-000-001"})
	address.AddContact(Contact{ID: 2, Name: "puge", Address: "Osake", Phone: "0120-000-002"})

	fmt.Println(address.GetContact(1))
	fmt.Println(address.GetContact(2))

	fmt.Println(address.ListContact())
}
