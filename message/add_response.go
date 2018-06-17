package message

import "fmt"

//
//        AddResponse ::= [APPLICATION 9] LDAPResult
func readAddResponse(bytes *Bytes) (ret AddResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagAddResponse)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAddResponse:\n%s", err.Error())}
		return
	}
	ret = AddResponse(res)
	return
}

//
//        AddResponse ::= [APPLICATION 9] LDAPResult
func (a AddResponse) write(bytes *Bytes) int {
	return LDAPResult(a).writeTagged(bytes, classApplication, TagAddResponse)
}

//
//        AddResponse ::= [APPLICATION 9] LDAPResult
func (a AddResponse) size() int {
	return LDAPResult(a).sizeTagged(TagAddResponse)
}
