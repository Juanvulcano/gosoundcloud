package gosoundcloud

import (
    "errors"
    "strconv"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

// For what I understand you cannot create nor delete a group via api. - to confirm
type Group struct {
    Id                  uint64
    Kind                string // to confirm if soundcloud respond with a kind for group types.
    Created_at          string "2009/06/18 15:46:46 +0000"
    Permalink           string "made-with-ableton-live"
    Name                string "Made with Ableton Live!"
    Short_description   string "tracks produced with Ableton Live! music software, no DJ mixes!"
    Description         string "send your tracks, no DJ mixes please!"
    Uri                 string "http://api.soundcloud.com/groups/3"
    Artwork_url         string "http://i1.sndcdn.com/artworks-000000481400-f0ynk3-large.jpg?142a848"
    Permalink_url       string "http://soundcloud.com/groups/made-with-ableton-live"
    Creator             User
    /* user contains based on api docs
    "creator": {
        "id": 1433,
        "permalink": "matas",
        "username": "matas",
        "uri": "http://api.soundcloud.com/users/1433",
        "permalink_url": "http://soundcloud.com/matas",
        "avatar_url": "http://i1.sndcdn.com/avatars-000001548772-zay6dy-large.jpg?142a848"
    }*/
}

func NewGroup() *Group {
    return &Group{Kind: "group"}
}

func (g Group) GetId() uint64 {
    return g.Id
}

func (g Group) GetKind() string {
    return g.Kind
}

func (g Group) IsNew() bool {
    if g.Id > 0 {
        return false
    }
    return true
}

func (g *Group) Save(s *SoundcloudApi) error {
    if !g.IsNew() {
        return errors.New("Group is not new, cannot be saved!")
    }
    // save group
    return nil
}

func (g *Group) Delete(s *SoundcloudApi) error {
    if g.IsNew() {
        return errors.New("Group is new, cannot be deleted!")
    }

    url := "/groups/" + strconv.FormatUint(g.Id, 10)
    _, err := s.Delete(url)
    return err
}

func getGroups(s *SoundcloudApi, p UrlParams) ([]*Group, error) {
    url := "/groups"
    resp, err := s.Get(url, p)
    var slice []*Group
    if err = processGroupResponses(resp, err, &slice); err != nil {
        return nil, err
    }
    return slice, err
}

func (g *Group) getModerators(s *SoundcloudApi, p UrlParams) ([]*User, error) {
    url := g.Uri + "/moderators"
    resp, err := s.Get(url, p)
    var slice []*User
    if err = processGroupResponses(resp, err, &slice); err != nil {
        return nil, err
    }
    return slice, err
}

func (g *Group) getMembers(s *SoundcloudApi, p UrlParams) ([]*User, error) {
    url := g.Uri + "/members"
    resp, err := s.Get(url, p)
    var slice []*User
    if err = processGroupResponses(resp, err, &slice); err != nil {
        return nil, err
    }
    return slice, err
}

func (g *Group) getContributors(s *SoundcloudApi, p UrlParams) ([]*User, error) {
    url := g.Uri + "/contributors"
    resp, err := s.Get(url, p)
    var slice []*User
    if err = processGroupResponses(resp, err, &slice); err != nil {
        return nil, err
    }
    return slice, err
}

func (g *Group) getUsers(s *SoundcloudApi, p UrlParams) ([]*User, error) {
    url := g.Uri + "/users"
    resp, err := s.Get(url, p)
    var slice []*User
    if err = processGroupResponses(resp, err, &slice); err != nil {
        return nil, err
    }
    return slice, err
}

func (g *Group) getTracks(s *SoundcloudApi, p UrlParams) ([]*Track, error) {
    url := g.Uri + "/tracks"
    resp, err := s.Get(url, p)
    var slice []*Track
    if err = processGroupResponses(resp, err, &slice); err != nil {
        return nil, err
    }
    return slice, err
}

func (g *Group) getPendingTracks(s *SoundcloudApi, p UrlParams) ([]*Track, error) {
    url := g.Uri + "/pending_tracks"
    resp, err := s.Get(url, p)
    var slice []*Track
    if err = processGroupResponses(resp, err, &slice); err != nil {
        return nil, err
    }
    return slice, err
}

// should be redundant with GetTrack unless the track resouce have adicional data here - to confirm
//func (s *SoundcloudApi) GetGroupPendingTrack(g *Group, id uint64) (*Track, error) {
//}

func (g *Group) getContributions(s *SoundcloudApi, p UrlParams) ([]*Track, error) {
    url := g.Uri + "/contributions"
    resp, err := s.Get(url, p)
    var slice []*Track
    if err = processGroupResponses(resp, err, &slice); err != nil {
        return nil, err
    }
    return slice, err
}

// should be redundant with GetTrack unless the track resouce have adicional data here - to confir
//func (s *SoundcloudApi) GetGroupContributionsTrack(g *Group) ([]*Track, error) {
//}

func (g Group) MarshalJSON() ([]byte, error) {
    j := map[string]map[string]interface{}{
        "group": {
            "name": g.Name,
            "short_description": g.Short_description,
            "description": g.Description,
            "auto_approve": false,
        },
    }
    return json.Marshal(j)
}

func processGroupResponses(resp *http.Response, err error, slice interface{}) (error) {
    if err != nil {
        return err
    }
    data, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        return err
    }
    if err = json.Unmarshal(data, slice); err != nil {
        return err
    }
    return err
}
