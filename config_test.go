package config

import (
    "fmt"
    "testing"
)

func TestParse(t *testing.T) {
    cfgs, err := Load("./test_cfg.ini")
    if err != nil {
        t.Error(err)
    }

    if cfgs.Has("no_exist_key") {
        t.Errorf("no exist key error")
    }

    if v := cfgs.Get("dbhost"); len(v) > 0 {
        t.Errorf("key dbhost should be empty")
    }

    if v, _ := cfgs.GetBool("usecache"); v {
        t.Errorf("key usecache should be false")
    }

    if v, _ := cfgs.GetInt("mgrPort"); v != 8866 {
        t.Errorf("key mgrPort should be int and equal 8866")
    }

    m := cfgs.Map()
    fmt.Println(m)
}
