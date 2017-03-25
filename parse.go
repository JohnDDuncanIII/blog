package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// user defined variables (this will be dynamic soon)
var path string = "file:///C:/Users/duncjo01/Documents/home/duncjo01/workspace_Go/src/github.com/JohnDDuncanIII/faces/"
var title = "John's Weblog"
var date_format = "Monday, January 2 2006 at 3:04pm"

func main() {
	/*now := time.Now()
	epoch := now.Unix()
	fmt.Println("Now: ", now)
	fmt.Println("Epoch(Unix) Time: ", epoch)
	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))
	fmt.Println("Day: ", now.Day())
	fmt.Println("Month: ", int(now.Month()))
	fmt.Println("Year: ", now.Year())
	fmt.Println("Hour: ", now.Hour())
	fmt.Println("Minute: ", now.Minute())
	fmt.Println(time.Unix(now.Unix(), 0).Format("Mon, Jan 2, 2006 at 3:04pm"))*/

	// holds the html for individual posts (archive)
	days := ""

	// when we come across a new month/year, add it to the map
	var extant_months map[string]bool
	extant_months = make(map[string]bool)

	// holds the html for months (archive)
	months := ""

	// when we come across a new year, add it to the map
	var extant_years map[string]bool
	extant_years = make(map[string]bool)

	// holds the html for years (archive)
	years := ""

	index_archive := ""

	entry_num := 0
	filename := "entries/" + strconv.Itoa(entry_num) + ".entry"
	_, e := os.Stat(filename);
	for e == nil {
		// extract specficic data from entry file
		postNum, name, subject, datetime, archive_name, content, more_content, comments, num_comments := parse_entries(filename)

		// convert datetime back to time object
		t, err := time.Parse(date_format, datetime)
		if err != nil {
			fmt.Println(err)
		}
		month := t.Month().String()
		year := strconv.Itoa(t.Year())

		// add links to the previous and next existing posts
		prev_path := "entries/" + strconv.Itoa(entry_num-1) + ".entry"
		next_path := "entries/" + strconv.Itoa(entry_num+1) + ".entry"
		prev_post := ""
		next_post := ""

		// check to see if the files actually exist on disk
		if _, err := os.Stat(prev_path); err == nil {
			prev_postNum, _, prev_subject, _, _, _, _, _, _ := parse_entries(prev_path)
			prev_post = `[<a href="` + path + `entries/` + prev_postNum + `.html">Previous entry: "` + prev_subject + `"</a>]`
		}
		if _, err := os.Stat(next_path); err == nil {
			next_postNum, _, next_subject, _, _, _, _, _, _ := parse_entries(next_path)
			next_post = `[<a href="` + path + `entries/` + next_postNum + `.html">Next entry: "` + next_subject + `"</a>]`
		}

		// html for individual posts (archive)
		days += `<a href="` + path +`entries/` +postNum +`.html">` + subject+`: ` + datetime + `</a><br>`

		// html for months (archive)
		if(!extant_months[month+"/"+year]) {
			extant_months[month+"/"+year] = true
			months += `<a href="`+ year + "/" + month +`.html">`+ month + " " + year +`</a><br>`

			months_archive, _ := parse_archive_write(month, year)
			index_archive += months_archive

			archive_month:= generate_archive_month(months_archive, month, year)
			if _, err := os.Stat("entries/" + year + "/"); os.IsNotExist(err) {
				os.Mkdir("entries/" + year, os.ModePerm)
			}

			archive_month_write := ioutil.WriteFile("entries/" + year + "/" + month + ".html", []byte(archive_month), 0644)

			if archive_month_write != nil {
				panic(archive_month_write)
			}
		}

		// html for years (archive)
		if(!extant_years[year]) {
			extant_years[year] = true
			years += `<a href="` + year + `/index.html">` + year + `</a><br>`

			_, years_archive := parse_archive_write("-1", year)

			archive_year := generate_archive_year(years_archive, year)

			if _, err := os.Stat("entries/" + year+"/"); os.IsNotExist(err) {
				os.Mkdir("entries/" + year, os.ModePerm)
			}

			archive_year_write := ioutil.WriteFile("entries/" + year + "/index.html", []byte(archive_year), 0644)

			if archive_year_write != nil {
				panic(archive_year_write)
			}
		}

		// write out individual entry to disk
		entries := generate_posts(subject, archive_name, month, year, prev_post, next_post, name, datetime, postNum, content, more_content, num_comments, comments)
		post_write := ioutil.WriteFile("entries/"+strconv.Itoa(entry_num)+".html", []byte(entries), 0644)
		if post_write != nil {
			panic(post_write)
		}

		// incr the entry # that we are parsing
		entry_num++

		// make sure the next file exists
		filename = "entries/" + strconv.Itoa(entry_num) + ".entry"
		_, e = os.Stat(filename);
	}

	// write out main entries log archive
	archive := generate_archive(days, months, years)
	archive_write := ioutil.WriteFile("entries/index.html", []byte(archive), 0644)
	if archive_write != nil {
		panic(archive_write)
	}

	index_html := generate_index(index_archive)
	archive_index_write := ioutil.WriteFile("index.html", []byte(index_html), 0644)
	if archive_index_write != nil {
		panic(archive_index_write)
	}
	fmt.Println("Success!")
}

func parse_archive_write(m string, y string) (string, string) {
	entry_num := 0
	month_r := ""
	year_r := ""
	filename := "entries/" + strconv.Itoa(entry_num) + ".entry"
	day_map := make(map[string]string)

	month_map := make(map[string]string)
	_, e := os.Stat(filename);

	for e == nil {
		day_html := ""
		month_html := ""
		postNum, name, subject, datetime, _, content, more_content, _, num_comments := parse_entries(filename)
		t, err:= time.Parse(date_format, datetime)
		if err != nil {
			fmt.Println(err)
		}
		day := strconv.Itoa(t.Day())
		month := t.Month().String()
		year := strconv.Itoa(t.Year())
		if y == year {
			if(m == month) {
				if(day_map[day] == "") {
					day_map[day] += `<div class="post">
<span class="raised">` + month + " " + day + " " + year + `</span>`
				}

				day_html += `<div class="content">
<h2>` + subject + `</h2>
<p>
` + content + `
</p>
<hr width="50%">
<p style="margin:0">
` + more_content + `
</p>
<div class="info">` + name + ` on ` + datetime + ` [<a href="` + path + `entries/` + postNum + `.html" title="` + month + "/" + day + "/" + year + `: ` + subject + `">link</a>][<a href="` + path + `entries/` + postNum + `.html#comments">`+num_comments+` Comments</a>]</div>
</div><hr>
`
				day_map[day] += day_html
			}

			if(month_map[month] == "") {
				month_map[month] += `<div class="post">
<span class="raised">` + month + " " + year + `</span>`
			}

			month_html += `<div class="content">
<h2>` + subject + `</h2>
<p>
` + content + `
</p>
<hr width="50%">
<p style="margin:0">
` + more_content + `
</p>
<div class="info">` + name + ` on ` + datetime + ` [<a href="` + path + `entries/` + postNum + `.html" title="` + month + "/" + day + "/" + year + `: ` + subject + `">link</a>][<a href="` + path + `entries/` + postNum + `.html#comments">` + num_comments + ` Comments</a>]</div>
</div><hr>
`
			month_map[month] += month_html
		}

		entry_num++
		filename = "entries/" + strconv.Itoa(entry_num) + ".entry"
		_, e = os.Stat(filename);
	}

	for k, _ := range day_map { // order doesn't matter here...
		day_map[k] += `</div><!-- end post -->`
	}

	for k, _ := range month_map { // order doesn't matter here...
		month_map[k] += `</div><!-- end post -->`
	}

	m_vals := []string{}
	for _, v := range day_map { m_vals = append(m_vals,v) }
	sort.Strings(m_vals)
	for _, k := range m_vals {
		month_r += k
	}

	y_vals := []string{}
	for _, v := range month_map { y_vals = append(y_vals,v) }
	sort.Strings(y_vals)
	for _, k := range y_vals {
		year_r += k
	}

	return month_r, year_r;
}

func parse_entries(filename string) (string, string, string, string, string, string, string, []string, string) {
	file, e := os.Open(filename)
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	metadata := strings.Split(scanner.Text(), "¦")

	postNum := metadata[0]
	name := metadata[1]
	subject := metadata[2]
	epoch, e := strconv.ParseInt(metadata[3], 10, 64)
	if e != nil {
		panic(e)
	}
	datetime := time.Unix(epoch, 0).Format(date_format)
	archive_name := strconv.Itoa(time.Unix(epoch, 0).Year()) + "/" +
		time.Unix(epoch, 0).Month().String()

	scanner.Scan()
	//ip  := scanner.Text()

	scanner.Scan()
	content  := scanner.Text()
	content = to_markdown(content)
	content = parse_emoticons(content)

	scanner.Scan()
	more_content  := scanner.Text()
	more_content = to_markdown(more_content)
	more_content = parse_emoticons(more_content)

	var comments []string
	for scanner.Scan() {
		comments = append(comments, scanner.Text())
	}

	num_comments := strconv.Itoa(len(comments))

	return postNum, name, subject, datetime, archive_name, content, more_content, comments, num_comments
}

func parse_comments(c []string) string {
	var toReturn string
	//var c_arr []string
	//ct := -1 // counter for overall comments
	//sub_ct := 1 // counter for replies
	//isPrevChild := false // determine whether we need to close a div or not

	for index := range c {
		cmt_splt := strings.Split(c[index],"¦")
		cmt_name := cmt_splt[0]
		//cmt_ip := cmt_splt[1]
		cmt_email := cmt_splt[2]
		cmt_hmpg := cmt_splt[3]
		cmt_epoch, err := strconv.ParseInt(cmt_splt[4], 10, 64)
		if err != nil {
			panic(err)
		}
		cmt_datetime := time.Unix(cmt_epoch, 0).Format(date_format)
		cmt_content := cmt_splt[5]
		cmt_content = parse_emoticons(cmt_content)
		cmt_face := cmt_splt[6]
		cmt_xface := cmt_splt[7]
		//cmt_deeper := cmt_splt[8]
		deeper := ""

		/*if(cmt_deeper != "") {
			isPrevChild = true
			if(len(cmt_deeper) < sub_ct) {
				for i:= 1; i < sub_ct; i++ {
					c_arr[ct] += `</div>
`
				}

				sub_ct = len(cmt_deeper)
			}

			deeper = " deeper"
			sub_ct++
		} else {
			// close out any child divs that may exist at the end of parent
			for i:= 1; i < sub_ct; i++ {
				c_arr[ct] += `</div>
`
				if(i == (sub_ct - 1)){
					sub_ct = 1
					c_arr[ct] += `</div>
`
					//reset = true
				}
			}
			// close out standard divs that do not contain any children
			if(ct > -1) {
				if(!isPrevChild) {
					c_arr[ct] += `</div>
`
				}
			}
			c_arr = append(c_arr, "")
			ct++
			isPrevChild = false
		}*/

		toReturn += `<div id="comment` + strconv.Itoa(index) + `" class="commentBox` + deeper + `"><div id="facesBox` + strconv.Itoa(index) + `" class="facesBox"><div id="picons` + strconv.Itoa(index) + `" class="picons">` + strings.Replace(strings.Trim(fmt.Sprint(search_picons(cmt_email)), "[]"), "> ", ">", -1) + `</div></div><p style="margin-top:0" align="left"> on ` + cmt_datetime + `, <a href="` + cmt_hmpg + `" target="_new">` + cmt_name + `</a> [<a href="mailto:` + cmt_email + `" rel="nofollow">e-mail</a>] said
</p><p align="justify">
` + cmt_content + `
</p><script>doGravatar("` + cmt_email + `");`
		if cmt_face != "" {
			toReturn += "doFace(\"" + cmt_face + "\");"
		}
		if cmt_xface != "" {
			toReturn += "doXFace(\"" + cmt_xface + "\");"
		}
		toReturn += `gCount++;</script><hr></div>
`
	}
	/*for _,element := range c_arr {
		toReturn += element
	}*/

	return toReturn
}


// TODO: move this to faces package
func search_picons(s string) []string {
	var pBox []string
	if(s=="") {
		pImg := `<img class="face" src="face/picons/misc/MISC/noface/face.gif" title="picon">`
		pBox = append(pBox, pImg)
	} else {
		atSign := strings.Index(s, "@");
		mfPiconDatabases := [4]string{"domains/", "users/", "misc", "usenix/"}
		count := 0
		// if we have a valid email address
		if (atSign != -1) {
			host := s[atSign + 1:len(s)]
			user := s[0:atSign]
			host_pieces := strings.Split(host, ".")

			pDef := `<img class="face" src="` + path + `face/picons/unknown/` + host_pieces[len(host_pieces)-1] + `/unknown/face.gif" title="` + host_pieces[len(host_pieces)-1] + `">`
			pBox = append(pBox, pDef)

			for i := range mfPiconDatabases {
				p_path := "face/picons/" + mfPiconDatabases[i]; // they are stored in $PROFILEPATH$/messagefaces/picons/ by default
				if mfPiconDatabases[i] == "misc/" {
					p_path += "MISC/"
				} // special case MISC

				// get number of database folders (probably six, but could theoretically change)
				var l = len(host_pieces)-1
				// we will check to see if we have a match at EACH depth,
				//     so keep a cloned version w/o the 'unknown/face.gif' portion
				for l >= 0 { // loop through however many pieces we have of the host
					p_path += host_pieces[l] + "/" // add that portion of the host (ex: 'edu' or 'gettysburg' or 'cs')
					clonedLocal := p_path
					if mfPiconDatabases[i] == "users/" {
						p_path += user + "/"
					} else {
						p_path += "unknown/"
					}
					p_path += "face.gif"
					if _, err := os.Stat(p_path); err == nil {
						if(count==0) {
							pBox[0] = `<img class="face" src="` + path + p_path + `"`
							if strings.Contains(p_path,"users") {
								pBox[0] += ` title="` + host_pieces[len(host_pieces)-1] + `">`
							} else {
								pBox[0] += ` title="` + host_pieces[l] + `">`
							}
						} else {
							pImg := `<img class="face" src="` + path + p_path + `"`
							if strings.Contains(p_path, "users") {
								pImg += ` title="` + user + `">`
							} else {
								pImg += ` title="` + host_pieces[l] + `">`
							}
							pBox = append(pBox, pImg)
						}
						count++;
					}
					p_path = clonedLocal;
					l--;
				}
			}
		}
	}
	return pBox
}


// TODO: move this to manip package
func to_markdown(s string) string {
	return strings.Replace(s, "|*|", "<br>", -1)
}


func parse_emoticons(s string) string {
	e_path := "<img src=" + path + "img/emoticons/"
	s = strings.Replace(s,":angry:",e_path + "angry.gif>",-1)
	s = strings.Replace(s,">:(",e_path + "angry.gif>",-1)
	s = strings.Replace(s,":laugh:",e_path + "laugh.gif>",-1)
	s = strings.Replace(s,":DD",e_path + "laugh.gif>",-1)
	s = strings.Replace(s,":yell:",e_path + "yell.gif>",-1)
	s = strings.Replace(s,">:O",e_path + "yell.gif>",-1)
	s = strings.Replace(s,":innocent:",e_path + "innocent.gif>",-1)
	s = strings.Replace(s,"O:)",e_path + "innocent.gif>",-1)
	s = strings.Replace(s,":satisfied:",e_path + "satisfied.gif>",-1)
	s = strings.Replace(s,"/:D",e_path + "satisfied.gif>",-1)
	s = strings.Replace(s,":)",e_path + "smile.gif>",-1)
	s = strings.Replace(s,":O",e_path + "shocked.gif>",-1)
	s = strings.Replace(s,":(",e_path + "sad.gif>",-1)
	s = strings.Replace(s,":D",e_path + "biggrin.gif>",-1)
	s = strings.Replace(s,":P",e_path + "tongue.gif>",-1)
	s = strings.Replace(s,";)",e_path + "wink.gif>",-1)
	s = strings.Replace(s,":blush",e_path + "blush.gif>",-1)
	s = strings.Replace(s,":\")",e_path + "blush.gif>",-1)
	s = strings.Replace(s,":confused:",e_path + "confused.gif>",-1)
	s = strings.Replace(s,":S",e_path + "confused.gif>",-1)
	s = strings.Replace(s,":cool:",e_path + "cool.gif>",-1)
	s = strings.Replace(s,"B)",e_path + "cool.gif>",-1)
	s = strings.Replace(s,":crazy:",e_path + "crazy.gif>",-1)
	s = strings.Replace(s,":cry:",e_path + "cry.gif>",-1)
	s = strings.Replace(s,":~(",e_path + "cry.gif>",-1)
	s = strings.Replace(s,":doze",e_path + "doze.gif>",-1)
	s = strings.Replace(s,":?",e_path + "doze.gif>",-1)
	s = strings.Replace(s,":hehe:",e_path + "hehe.gif>",-1)
	s = strings.Replace(s,"XD",e_path + "hehe.gif>",-1)
	s = strings.Replace(s,":plain:",e_path + "plain.gif>",-1)
	s = strings.Replace(s,":|",e_path + "plain.gif>",-1)
	s = strings.Replace(s,":rolleyes:",e_path + "rolleyes.gif>",-1)
	s = strings.Replace(s,"9_9",e_path + "rolleyes.gif>",-1)
	s = strings.Replace(s,":dizzy:",e_path + "crazy.gif>",-1)
	s = strings.Replace(s,"o_O",e_path + "crazy.gif>",-1)
	s = strings.Replace(s,":money:",e_path + "money.gif>",-1)
	s = strings.Replace(s,":$",e_path + "money.gif>",-1)
	s = strings.Replace(s,":sealed:",e_path + "sealed.gif>",-1)
	s = strings.Replace(s,":X",e_path + "sealed.gif>",-1)
	s = strings.Replace(s,":eek:",e_path + "eek.gif>",-1)
	s = strings.Replace(s,"O_O",e_path + "eek.gif>",-1)
	s = strings.Replace(s,":kiss:",e_path + "kiss.gif>",-1)
	s = strings.Replace(s,":*",e_path + "kiss.gif>",-1)

	return s
}

// TODO: move this to html/tmpl package
// TODO: actually turn this into a separate template file
func generate_posts(subject string, archive_name string, month string, year string, prev_post string, next_post string, name string, datetime string, postNum string, content string, more_content string, num_comments string, comments []string) string {
	toReturn := `<!DOCTYPE HTML>
<head><title>` + title + `: ` + subject + `</title>
<meta charset="UTF-8">
<meta name="generator" content="DarkMatter 1.8.3">
<link rel="stylesheet" href="` + path + `css/gm.css">
<link rel="stylesheet" href="` + path + `css/face.css">
</head>
<body>
<div id="frame">
<h1 id="header" class="header">` + title + `</h1>
<!-- <div id="contentright">
{sidebar}
</div>-->
<div class="path"><a href="` + path + `index.html" title="back to frontpage">Home</a> &raquo; <a href="` + path + `entries/index.html" title="weblog entries">Entries</a> &raquo; <a href="` + path + `entries/` + year + `/index.html" title="archive of ` + year + `">` + year + `</a> &raquo; <a href="` + path + `entries/` + year + "/" + month + `.html" title="archive of ` + month + `">` + month + `</a> &raquo; ` + subject + `</div>
<div class="direction">
` + prev_post + " " + next_post + `
</div>
<div id="contentcenter">
<div class="post">

<h2 class="h2_full">` + subject + `</h2>
<div class="info info_archive">` + name + ` on ` + datetime + ` [<a href="` + path + `entries/` + postNum + `.html" title="` + subject + `">permalink</a>]
</div>
<p>
` + content + `
</p>
<hr>
<p style="margin:0">
` + more_content + `
</p>
</div>

<script src="` + path + `face/xface.js"></script>
<script src="` + path + `face/md5-impl.js"></script>
<script src="` + path + `face/main.js"></script>
<script>gCount = 0;</script>
<div id="comments">
<a name="comments"> </a>
<p align="center">
<strong>Replies: ` + num_comments + ` Comments</strong>
</p>
` + parse_comments(comments) + `
<!-- commentsform code begin -->
<div align="center">
<form id="new_comment_box" action="` + path + `cgi-bin/gm-comments.cgi#comments" method="post" name="newcomment" display="block">

<input name="newcommententrynumber" type="hidden" value="2">
<span style="font-weight:bold;">New Comment</span>

<input name="newcommentauthor" placeholder="name" type="text" class="text">
<input name="newcommentemail" placeholder="email" type="text" class="text">
<input name="newcommentxface" placeholder="x-face" type="text" class="text">
<input name="newcommentface" placeholder="face" type="text" class="text">
<input name="newcommenthomepage" placeholder="homepage" type="text" class="text">

<div id="input_box">
<div id="emoticons">
Smilies:
<div>
<img onclick="commentEmoticon(':)')" src="` + path + `img/emoticons/smile.gif" alt="smile">
<img onclick="commentEmoticon(':O')" src="` + path + `img/emoticons/shocked.gif" alt="shocked">
<img onclick="commentEmoticon(':(')" src="` + path + `img/emoticons/sad.gif" alt="sad">
</div>

<div>
<img onclick="commentEmoticon(':D')" src="` + path + `img/emoticons/biggrin.gif" alt="big grin">
<img onclick="commentEmoticon(':P')" src="` + path + `img/emoticons/tongue.gif" alt="razz">
<img onclick="commentEmoticon(';)')" src="` + path + `img/emoticons/wink.gif" alt="*wink wink* hey baby">
</div>

<div>
<img onclick="commentEmoticon(':angry:')" src="` + path + `img/emoticons/angry.gif" alt="angry, grr">
<img onclick="commentEmoticon(':blush:')" src="` + path + `img/emoticons/blush.gif" alt="blush">
<img onclick="commentEmoticon(':confused:')" src="` + path + `img/emoticons/confused.gif" alt="confused">
</div>

<div>
<img onclick="commentEmoticon(':cool:')" src="` + path + `img/emoticons/cool.gif" alt="cool">
<img onclick="commentEmoticon(':crazy:')" src="` + path + `img/emoticons/crazy.gif" alt="crazy">
<img onclick="commentEmoticon(':cry:')" src="` + path + `img/emoticons/cry.gif" alt="cry">
</div>

<div>
<img onclick="commentEmoticon(':doze:')" src="` + path + `img/emoticons/doze.gif" alt="sleepy">
<img onclick="commentEmoticon(':hehe:')" src="` + path + `img/emoticons/hehe.gif" alt="hehe">
<img onclick="commentEmoticon(':laugh:')" src="` + path + `img/emoticons/laugh.gif" alt="LOL">
</div>

<div>
<img onclick="commentEmoticon(':plain:')" src="` + path + `img/emoticons/plain.gif" alt="plain jane">
<img onclick="commentEmoticon(':rolleyes:')" src="` + path + `img/emoticons/rolleyes.gif" alt="rolls eyes">
<img onclick="commentEmoticon(':satisfied:')" src="` + path + `img/emoticons/satisfied.gif" alt="satisfied">
</div></div>

<textarea name="newcommentbody"></textarea>
</div>
<div>
<input id="bakecookie" name="bakecookie" type="checkbox">Save Info?
<label for="bakecookie">
<span></span>
</label>
</div>
<input type="reset" value="Reset" class="button">
<input name="gmpostpreview" type="submit" value="Preview" class="button">
<input type="submit" value="Submit" class="button"  onClick="javascript:setGMlocalStorage()">
</form>
</div>
<script>
function commentEmoticon(code)
{
	var cache = document.newcomment.newcommentbody.value;
	document.newcomment.newcommentbody.value = cache + " " + code;
}
document.newcomment.newcommentauthor.value = localStorage.getItem("gmcmtauth");
document.newcomment.newcommentemail.value = localStorage.getItem("gmcmtmail");
document.newcomment.newcommenthomepage.value = localStorage.getItem("gmcmthome");
function setGMlocalStorage(){
if(document.getElementById("bakecookie").checked){
	localStorage.setItem("gmcmtauth", document.newcomment.newcommentauthor.value);
	localStorage.setItem("gmcmtmail", document.newcomment.newcommentemail.value);
	localStorage.setItem("gmcmthome", document.newcomment.newcommenthomepage.value);
}else{ localStorage.removeItem("gmcmtauth");localStorage.removeItem("gmcmtmail");localStorage.removeItem("gmcmthome"); }}</script>
<!-- commentsform code end -->
</div>
</div><div id="contentsidebar"><div><a href="` + path + `index.html">Home</a><br>
<a href="` + path + `entries/index.html">Entries</a><br>

<a href="#">Fake Link One</a><br>
<a href="#">Fake Link Two</a><br>
<a href="#">Fake Link Three</a><br><br>

<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_blank">DarkMatter Source</a></div>
<hr>
<!-- calendar code begin -->
<!-- calendar code end -->
<hr>
<!-- searchform code begin -->
<div class="searchform">
<form action="` + path + `cgi-bin/gm-comments.cgi" method="post"><div><input type="text" name="gmsearch" class="text"></div>
<div><input type="submit" value="Search" class="button"></div></form></div>
<!-- searchform code end -->
<hr>
<div align="center">
<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_top"><img src="` + path + `img/dm_1.8.3.gif" alt="Powered By Greymatter"></a><a href="http://validator.w3.org/check/referer"><img src="` + path + `img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div>
</div>
<script src="` + path + `js/scroll.js"></script>
</body>
`
	return toReturn
}

func generate_archive (days string, months string, years string) string {
	archive := `<!DOCTYPE HTML>
<html><head><title>` + title + `</title>
<meta charset="UTF-8">
<meta name="generator" content="DarkMatter 1.8.3">
<link rel="stylesheet" href="` + path + `css/gm.css">
</head>
<body>
<div id="frame">
<h1 id="header" class="header"> ` + title + ` </h1>

<!-- <div id="contentright">
{sidebar}
</div>-->
<div class="path"><a href="` + path + `index.html" title="back to frontpage">Home</a> &raquo; Entries</div>
<div id="contentcenter">
<div class="content">
<h1>Years</h1>
<p>` + years + `</p>
</div>
<div class="content">
<h1>Months</h1>
<p>` + months + `</p>
</div>
<div class="content">
<h1>Entries</h1>
<p>` + days + `</p>
</div>
</div><div id="contentsidebar"><div><a href="` + path + `index.html">Home</a><br>
<a href="` + path + `entries/index.html">Entries</a><br>

<a href="#">Fake Link One</a><br>
<a href="#">Fake Link Two</a><br>
<a href="#">Fake Link Three</a><br><br>

<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_blank">DarkMatter Source</a></div>
<hr>
<!-- calendar code begin -->
<!-- calendar code end -->
<hr>
<!-- searchform code begin -->
<div class="searchform">
<form action="` + path + `cgi-bin/gm-comments.cgi" method="post"><div><input type="text" name="gmsearch" class="text"></div>
<div><input type="submit" value="Search" class="button"></div></form></div>
<!-- searchform code end -->
<hr>
<div align="center">
<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_top"><img src="` + path + `img/dm_1.8.3.gif" alt="Powered By Greymatter"></a><a href="http://validator.w3.org/check/referer"><img src="` + path + `img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div><!-- https://github.com/JohnDDuncanIII/DarkMatter/-->
</div>
<script src="` + path + `js/scroll.js"></script>
</body>`

	return archive
}


func generate_archive_month(months_archive string, month string, year string) string {
	archive_month := `<!DOCTYPE HTML>
<head><title>` + title + `</title>
<meta charset="UTF-8">
<meta name="generator" content="DarkMatter 1.8.3">
<link rel="stylesheet" href="` + path + `css/gm.css">
</head>
<body>
<div id="frame">
<h1 id="header" class="header"> ` + title + ` </h1>
<!-- <div id="contentright">
{sidebar}
</div>-->
<div class="path"><a href="` + path + `index.html" title="back to frontpage">Home</a> &raquo; <a href="` + path + `entries/index.html" title="weblog entries">Entries</a> &raquo; <a href="` + path + `entries/` + year + `/index.html" title="archive of ` + year + `">` + year + `</a> &raquo; ` + month + `</div>
<div id="contentcenter">
` + months_archive + `
</div><div id="contentsidebar"><div><a href="` + path + `index.html">Home</a><br>
<a href="` + path + `entries/index.html">Entries</a><br>

<a href="#">Fake Link One</a><br>
<a href="#">Fake Link Two</a><br>
<a href="#">Fake Link Three</a><br><br>

<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_blank">DarkMatter Source</a></div>
<hr>
<!-- calendar code begin -->
<!-- calendar code end -->
<hr>
<!-- searchform code begin -->
<div class="searchform">
<form action="` + path + `cgi-bin/gm-comments.cgi" method="post"><div><input type="text" name="gmsearch" class="text"></div>
<div><input type="submit" value="Search" class="button"></div></form></div>
<!-- searchform code end -->
<hr>
<div align="center">
<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_top"><img src="` + path + `img/dm_1.8.3.gif" alt="Powered By Greymatter"></a><a href="http://validator.w3.org/check/referer"><img src="` + path + `img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div><!-- https://github.com/JohnDDuncanIII/DarkMatter/-->
</div>
<script src="` + path + `js/scroll.js"></script>
</body>`
	return archive_month
}

func generate_archive_year(years_archive string, year string) string {
	archive_year := `<!DOCTYPE HTML>
<head><title>` + title + `</title>
<meta charset="UTF-8">
<meta name="generator" content="DarkMatter 1.8.3">
<link rel="stylesheet" href="` + path + `css/gm.css">
</head>
<body>
<div id="frame">
<h1 id="header" class="header"> ` + title + ` </h1>
<!-- <div id="contentright">
{sidebar}
</div>-->
<div class="path"><a href="` + path + `index.html" title="back to frontpage">Home</a> &raquo; <a href="` + path + `entries/index.html" title="weblog entries">Entries</a> &raquo; ` + year + `</div>
<div id="contentcenter">
` + years_archive + `
</div><div id="contentsidebar"><div><a href="` + path + `index.html">Home</a><br>
<a href="` + path + `entries/index.html">Entries</a><br>

<a href="#">Fake Link One</a><br>
<a href="#">Fake Link Two</a><br>
<a href="#">Fake Link Three</a><br><br>

<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_blank">DarkMatter Source</a></div>
<hr>
<!-- calendar code begin -->
<!-- calendar code end -->
<hr>
<!-- searchform code begin -->
<div class="searchform">
<form action="` + path + `cgi-bin/gm-comments.cgi" method="post"><div><input type="text" name="gmsearch" class="text"></div>
<div><input type="submit" value="Search" class="button"></div></form></div>
<!-- searchform code end -->
<hr>
<div align="center">
<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_top"><img src="` + path + `img/dm_1.8.3.gif" alt="Powered By Greymatter"></a><a href="http://validator.w3.org/check/referer"><img src="` + path + `img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div><!-- https://github.com/JohnDDuncanIII/DarkMatter/-->
</div>
<script src="` + path + `js/scroll.js"></script>
</body>`
	return archive_year
}


func generate_index (index_archive string) string {
	main_index := `<!DOCTYPE HTML>
<head><title>` + title + `</title>
<meta charset="UTF-8">
<meta name="generator" content="DarkMatter 1.8.3">
<link rel="stylesheet" href="` + path + `css/gm.css">
</head>
<body>
<div id="frame">
<h1 id="header" class="header"> ` + title + ` </h1>
<!--<div id="contentright">
{sidebar}
</div>-->
<div id="contentcenter">
` + index_archive + `
</div><div id="contentsidebar"><div><a href="` + path + `index.html">Home</a><br>
<a href="` + path + `entries/index.html">Entries</a><br>

<a href="#">Fake Link One</a><br>
<a href="#">Fake Link Two</a><br>
<a href="#">Fake Link Three</a><br><br>

<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_blank">DarkMatter Source</a></div>
<hr>
<!-- calendar code begin -->
<!-- calendar code end -->
<hr>
<!-- searchform code begin -->
<div class="searchform">
<form action="` + path + `cgi-bin/gm-comments.cgi" method="post"><div><input type="text" name="gmsearch" class="text"></div>
<div><input type="submit" value="Search" class="button"></div></form></div>
<!-- searchform code end -->
<hr>
<div align="center">
<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_top"><img src="` + path + `img/dm_1.8.3.gif" alt="Powered By Greymatter"></a><a href="http://validator.w3.org/check/referer"><img src="` + path + `img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div>
</div>
<script src="` + path + `js/scroll.js"></script>
</body>`
	return main_index
}
