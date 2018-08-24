package paracross

import "gitlab.33.cn/chain33/chain33/types"
import (
	dbm "gitlab.33.cn/chain33/chain33/common/db"
)

func (c *Paracross) ParacrossGetHeight(title string) (types.Message, error) {
	ret, err := getTitle(c.GetStateDB(), calcTitleKey(title))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Paracross) ParacrossListTitles() (types.Message, error) {
	return listLocalTitles(c.GetLocalDB())
}

func listLocalTitles(db dbm.KVDB) (types.Message, error) {
	prefix := calcLocalTitlePrefix()
	res, err := db.List(prefix, []byte(""), 0, 1)
	if err != nil {
		return nil, err
	}
	var resp types.RespParacrossTitles
	for _, r := range res {
		var st types.ReceiptParacrossDone
		err = types.Decode(r, &st)
		if err != nil {
			panic(err)
		}
		resp.Titles = append(resp.Titles, &st)
	}
	return &resp, nil
}

func loadLocalTitle(db dbm.KV, title string, height int64) (types.Message, error) {
	key := calcLocalTitleHeightKey(title, height)
	res, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	var resp types.ReceiptParacrossDone
	err = types.Decode(res, &resp)
	if err != nil {
		panic(err)
	}
	return &resp, nil
}
func (c *Paracross) ParacrossGetTitleHeight(title string, height int64) (types.Message, error) {
	return loadLocalTitle(c.GetLocalDB(), title, height)
}