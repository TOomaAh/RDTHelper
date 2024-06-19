package realdebridold

/*
type RealDebridAPI struct {
	db     *gorm.DB
	client *RealDebridClient
}

func NewRealDebridAPI(db *gorm.DB, client *RealDebridClient) *RealDebridAPI {
	return &RealDebridAPI{db, client}
}

func initHeader(c *gin.Context) map[string]string {
	db := c.MustGet("db").(*gorm.DB)

	var user model.User
	id, err := model.ExtractTokenID(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return nil
	}

	user.ID = id
	user.FindOne(db)

	rdtToken := user.RdtToken

	fmt.Println(rdtToken)

	//create header

	header := map[string]string{
		"Authorization": "Bearer " + rdtToken,
	}
	return header
}

func GetOne(c *gin.Context) *model.Torrent {
	return GetOneWithID(c, c.Param("id"))
}

var APIURL = "https://api.real-debrid.com/rest/1.0"

func GetOneWithID(c *gin.Context, id string) *model.Torrent {
	//Get One torrent from real debrid
	var torrent *model.Torrent

	header := initHeader(c)

	if header == nil {
		return nil
	}

	options := fast.Option{
		Method:      http.MethodGet,
		ContentType: fast.ContentTypeJSON,
		Headers:     header,
	}

	if err := fast.Fetch(APIURL+"/torrents/info/"+id, options).ToJSON(torrent); err != nil {
		log.Println(err)
		return nil
	}

	return torrent
}

func DeleteOne(c *gin.Context) {
	//Delete One torrent from real debrid

	header := initHeader(c)

	if header == nil {
		return
	}

	options := fast.Option{
		Method:      http.MethodDelete,
		ContentType: fast.ContentTypeJSON,
		Headers:     header,
	}

	if _, err := fast.Fetch(APIURL+"/torrents/delete/"+c.Param("id"), options).ToString(); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Torrent deleted"})
}

func AcceptOne(c *gin.Context) {
	//Accept One torrent from real debrid

	header := initHeader(c)

	if header == nil {
		return
	}

	body := map[string]string{
		"files": "all",
	}

	options := fast.Option{
		Method:      http.MethodPost,
		ContentType: "application/x-www-form-urlencoded",
		Headers:     header,
		Body:        body,
	}

	if _, err := fast.Fetch(APIURL+"/torrents/selectFiles/"+c.Param("id"), options).ToString(); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Torrent accepted"})
}

func Upload(c *gin.Context) *[]model.RdtUpload {
	//Upload torrent to real debrid
	header := initHeader(c)

	if header == nil {
		return nil
	}

	from, err := c.MultipartForm()

	if err != nil {
		return nil
	}

	files := from.File["file"]

	var rdtUpload []model.RdtUpload

	for _, file := range files {

		url := APIURL + "/torrents/addTorrent"
		method := "PUT"
		content, err := file.Open()
		if err != nil {
			return nil
		}
		defer content.Close()
		bytes, _ := ioutil.ReadAll(content)
		payload := strings.NewReader(string(bytes))

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			return nil
		}

		req.Header.Add("Authorization", header["Authorization"])
		req.Header.Add("Content-Type", "application/x-bittorrent")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		var uploadedFile model.RdtUpload
		//parse json body
		if err := json.Unmarshal(body, &uploadedFile); err != nil {
			return nil
		}
		rdtUpload = append(rdtUpload, uploadedFile)

	}

	return &rdtUpload
}

func Debrid(c *gin.Context) *real {
	//Debrid torrent from real debrid

	header := initHeader(c)

	if header == nil {
		return nil
	}

	//get link from body
	var linkRequest model.LinkRequest
	var link model.Link
	if err := c.ShouldBindJSON(&linkRequest); err != nil {
		log.Println(err)
		return nil
	}

	body := map[string]string{
		"link": linkRequest.Link,
	}
	options := fast.Option{
		Method:      http.MethodPost,
		ContentType: "application/x-www-form-urlencoded",
		Headers:     header,
		Body:        body,
	}

	if err := fast.Fetch(APIURL+"/unrestrict/link", options).ToJSON(&link); err != nil {
		log.Println(err)
		return nil
	}
	return &link
}
*/
