package json_parse

import "testing"

func TestParseJson(t *testing.T) {
	//parser_json()
	strjson := []byte(`{"optionsxxx":{"id":"3da3b1791aaa","conjunction":"and","children":[{"left":{"type":"field","field":"shezheng"},"op":"equal","right":"sss","id":"2c8d28ea195b","children":[{"id":"a64cdac5b0b4","op":"xxxx"}]},{"id":"435367d4c3da","conjunction":"and","children":[{"id":"a64cdac5b0b4"}]},{"id":"fa62a37b02cd","left":{"type":"field","field":"shifei"}}]}}`)
	jsonParse(strjson)
}


func TestTree_BuildTree(t *testing.T) {
	strjson := []byte(`{"optionsxxx":{"id":"3da3b1791aaa","conjunction":"and","children":[{"left":{"type":"field","field":"shezheng"},"op":"equal","right":"sss0","id":"000-000"},{"id":"000-001","conjunction":"and","children":[{"left":{"type":"field","field":"shezheng"},"op":"equal","right":"sss11","id":"111-000"},{"left":{"type":"field","field":"shezheng"},"op":"equal","right":"kkk11","id":"111-001"}]},{"left":{"type":"field","field":"shezheng"},"op":"equal","right":"sss01","id":"000-002"}]}}`)
	tree    := Tree{}
    tree.BuildTree(strjson)
	tree.bfsTree()
}