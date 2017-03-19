package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"strconv"
	"time"
)

// user defined variables (this will be dynamic soon)
var path string = "file:///C:/cygwin64/home/John/go_workspace/faces/"
var title = "John's Weblog"
var date_format = "Mon, Jan 2, 2006 at 3:04pm"

func main() {
	/*
	now := time.Now()
	epoch := now.Unix()

	fmt.Println("Now: ", now)
	fmt.Println("Epoch(Unix) Time: ", epoch)

	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	fmt.Println("Day: ", now.Day())
	fmt.Println("Month: ", int(now.Month()))
	fmt.Println("Year: ", now.Year())

	fmt.Println("Hour: ", now.Hour())
	fmt.Println("Minute: ", now.Minute())

	fmt.Println(time.Unix(now.Unix(), 0).Format("Mon, Jan 2, 2006 at 3:04pm"))
*/
	c := 0
	filename := "archives/" +strconv.Itoa(c)+".entry"
	_, e := os.Stat(filename);
	for e == nil {
		postNum, name, subject, datetime, archive_name, content, more_content, comments, num_comments := parse_entries(filename)
		// calculate the previous and next posts
		prev_path := "archives/" +strconv.Itoa(c-1)+".entry"
		next_path := "archives/" +strconv.Itoa(c+1)+".entry"
		prev_post := ""
		next_post := ""

		if _, err := os.Stat(prev_path); err == nil {
			prev_postNum, _, prev_subject, _, _, _, _, _, _ := parse_entries(prev_path)
			prev_post = `[<a href="`+path+`archives/`+prev_postNum+`.html">Previous entry: "`+prev_subject+`"</a>]`
		}
		if _, err := os.Stat(next_path); err == nil {
			next_postNum, _, next_subject, _, _, _, _, _, _ := parse_entries(next_path)
			next_post = `[<a href="`+path+`archives/`+next_postNum+`.html">Next entry: "`+next_subject+`"</a>]`
		}

		// parse the formatted string back into a time object
		t, err:= time.Parse(date_format, datetime)
		if err != nil {
			fmt.Println(err)
		}
		month := t.Month().String()
		year := strconv.Itoa(t.Year())

		tpl :=
`<!DOCTYPE HTML>
<head><title>`+title+`: `+subject+`</title>
<meta charset="UTF-8">
<meta name="generator" content="DarkMatter 1.8.3">
<link rel="stylesheet" href="`+path+`css/gm.css">
<link rel="stylesheet" href="`+path+`css/face.css">
</head>
<body>
<div id="frame">
<h1 id="header" class="header">`+title+`</h1>
<!-- <div id="contentright">
{sidebar}
</div>-->
<div class="path"><a href="`+path+`" title="back to frontpage">Home</a> &raquo; <a href="`+path+`archives/" title="weblog archives">Archives</a> &raquo; <a href="`+path+`archives/`+archive_name+`.html" title="archive of `+month+" " + year+`">`+month+" " + year+`</a> &raquo; `+subject+`</div>
<div class="direction">
`+prev_post+" "+next_post+`
</div>
<div id="contentcenter">
<div class="post">

<h2 class="h2_full">`+subject+`</h2>
<div class="info info_archive">`+name+` on `+datetime+` [<a href="`+path+`archives/`+postNum+`.html" title="`+subject+`">permalink</a>]
</div>
<p>
`+content+`
</p>
<hr>
<p style="margin:0">
`+more_content+`
</p>
</div>

<script src="`+path+`face/xface.js"></script>
<script src="`+path+`face/md5-call.js"></script>
<script src="`+path+`face/md5-impl.js"></script>
<script src="`+path+`face/main.js"></script>
<script>gCount = 0;</script>
<div id="comments">
<a name="comments"> </a>
<p align="center">
<strong>Replies: `+num_comments+` Comments</strong>
</p>
`+parse_comments(comments)+`
<!-- commentsform code begin -->
<div align="center">
<form id="new_comment_box" action="`+path+`cgi-bin/gm-comments.cgi#comments" method="post" name="newcomment" display="block">

<input name="newcommententrynumber" type="hidden" value="2" />
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
<img onclick="commentEmoticon(':)')" src="`+path+`emoticons/smile.gif" alt="smile">
<img onclick="commentEmoticon(':O')" src="`+path+`emoticons/shocked.gif" alt="shocked">
<img onclick="commentEmoticon(':(')" src="`+path+`emoticons/sad.gif" alt="sad">
</div>

<div>
<img onclick="commentEmoticon(':D')" src="`+path+`emoticons/biggrin.gif" alt="big grin">
<img onclick="commentEmoticon(':P')" src="`+path+`emoticons/tongue.gif" alt="razz">
<img onclick="commentEmoticon(';)')" src="`+path+`emoticons/wink.gif" alt="*wink wink* hey baby">
</div>

<div>
<img onclick="commentEmoticon(':angry:')" src="`+path+`emoticons/angry.gif" alt="angry, grr">
<img onclick="commentEmoticon(':blush:')" src="`+path+`emoticons/blush.gif" alt="blush">
<img onclick="commentEmoticon(':confused:')" src="`+path+`emoticons/confused.gif" alt="confused">
</div>

<div>
<img onclick="commentEmoticon(':cool:')" src="`+path+`emoticons/cool.gif" alt="cool">
<img onclick="commentEmoticon(':crazy:')" src="`+path+`emoticons/crazy.gif" alt="crazy">
<img onclick="commentEmoticon(':cry:')" src="`+path+`emoticons/cry.gif" alt="cry">
</div>

<div>
<img onclick="commentEmoticon(':doze:')" src="`+path+`emoticons/doze.gif" alt="sleepy">
<img onclick="commentEmoticon(':hehe:')" src="`+path+`emoticons/hehe.gif" alt="hehe">
<img onclick="commentEmoticon(':laugh:')" src="`+path+`emoticons/laugh.gif" alt="LOL">
</div>

<div>
<img onclick="commentEmoticon(':plain:')" src="`+path+`emoticons/plain.gif" alt="plain jane">
<img onclick="commentEmoticon(':rolleyes:')" src="`+path+`emoticons/rolleyes.gif" alt="rolls eyes">
<img onclick="commentEmoticon(':satisfied:')" src="`+path+`emoticons/satisfied.gif" alt="satisfied">
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
</div><div id="contentsidebar"><div><a href="`+path+`">Home</a><br />
<a href="`+path+`archives/">Archives</a><br />

<a href="#">Fake Link One</a><br />
<a href="#">Fake Link Two</a><br />
<a href="#">Fake Link Three</a><br /><br />

<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_blank">Greymatter Forums</a></div>
<hr>
<!-- calendar code begin -->
<!-- calendar code end -->
<hr>
<!-- searchform code begin -->
<div class="searchform">
<form action="`+path+`cgi-bin/gm-comments.cgi" method="post"><div><input type="text" name="gmsearch" class="text" /></div>
<div><input type="submit" value="Search" class="button" /></div></form></div>
<!-- searchform code end -->
<hr>
<div align="center">
<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_top"><img src="`+path+`img/dm_1.8.3.gif" alt="Powered By Greymatter" /></a><a href="http://validator.w3.org/check/referer"><img src="`+path+`img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div>
</div>
<script src="`+path+`js/scroll.js"></script>
</body>`
		//fmt.Println(tpl)
		// write the whole body at once
		post_write := ioutil.WriteFile("archives/"+strconv.Itoa(c)+".html", []byte(tpl), 0644)
		if post_write != nil {
			panic(post_write)
		}
		c++
		filename = "archives/" +strconv.Itoa(c)+".entry"
		_, e = os.Stat(filename);
	}
	archive:=
`<!DOCTYPE HTML>
<html><head><title>`+title+`</title>
<meta charset="UTF-8">
<link rel="stylesheet" href="`+path+`css/gm.css">
</head>
<body>
<div id="frame">
<h1 id="header" class="header"> `+title+` </h1>

<!-- <div id="contentright">
{sidebar}
</div>-->
<div class="path"><a href="`+path+`" title="back to frontpage">Home</a> &raquo; Archives</div>
<div id="contentcenter">
<div class="content">
<h1>Log Archives</h1>
<p>`+parse_months_archive()+`</p>
</div>

<div class="content">
<h1>Entries</h1>
<p>`+parse_entries_archive()+`</p>
</div>

</div><div id="contentsidebar"><div><a href="`+path+`">Home</a><br />
<a href="`+path+`archives/">Archives</a><br />

<a href="#">Fake Link One</a><br />
<a href="#">Fake Link Two</a><br />
<a href="#">Fake Link Three</a><br /><br />

<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_blank">Greymatter Forums</a></div>
<hr>
<!-- calendar code begin -->
<!-- calendar code end -->
<hr>
<!-- searchform code begin -->
<div class="searchform">
<form action="`+path+`cgi-bin/gm-comments.cgi" method="post"><div><input type="text" name="gmsearch" class="text" /></div>
<div><input type="submit" value="Search" class="button" /></div></form></div>
<!-- searchform code end -->
<hr>
<div align="center">
<a href="https://github.com/JohnDDuncanIII/DarkMatter/" target="_top"><img src="`+path+`img/dm_1.8.3.gif" alt="Powered By Greymatter" /></a><a href="http://validator.w3.org/check/referer"><img src="`+path+`img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div><!-- https://github.com/JohnDDuncanIII/DarkMatter/-->
</div>
<script src="`+path+`js/scroll.js"></script>
</body>`
	archive_write := ioutil.WriteFile("archives/index.html", []byte(archive), 0644)
	if archive_write != nil {
		panic(archive_write)
	}

	/*archive_month:=`<!DOCTYPE HTML>
<head><title>`+title+`</title>
<meta charset="UTF-8">
<link rel="stylesheet" href="`+path+`css/gm.css">
<link rel="stylesheet" href="`+path+`css/face.css">
</head>
<body>
<div id="frame">
<h1 id="header" class="header"> `+title+` </h1>
<!-- <div id="contentright">
{sidebar}
</div>-->
<div id="contentcenter">
`+parse_months_archive_write()+`
</div> <!-- end contentcenter -->
<div id="contentsidebar"><div><a href="`+path+`">Home</a><br>
<a href="`+path+`archives/">Archives</a><br>

<a href="#">Fake Link One</a><br>
<a href="#">Fake Link Two</a><br>
<a href="#">Fake Link Three</a><br><br>

<a href="http://greymatterforum.proboards.com/" target="_blank">Greymatter Forums</a></div>
<hr>
<!-- calendar code begin -->
<!-- calendar code end -->
<hr>
<!-- searchform code begin -->
<div class="searchform">
<form action="`+path+`cgi-bin/gm-comments.cgi" method="post"><div><input type="text" name="gmsearch" class="text"></div>
<div><input type="submit" value="Search" class="button"></div></form></div>
<!-- searchform code end -->
<hr>
<div align="center">
<a href="http://greymatterforum.proboards.com/" target="_top"><img src="`+path+`img/dm_1.8.3.gif" alt="Powered By Greymatter"></a><a href="http://validator.w3.org/check/referer"><img src="`+path+`img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div><!-- http://greymatterforum.proboards.com/-->
</div>
<script src="`+path+`js/scroll.js"></script>
</body>`*/
}


/*func parse_months_archive_write() {
<div class="post">
<span class="raised">Sunday, February 5th</span>
<div class="content">
<h2>#AltWoke 2</h2>
<p>
`+content+`
</p>
<hr>
<p style="margin:0">
`+more_content+`
</p>
<div class="info">Alice on 02.05.17 @ 05:52 PM CST [<a href="`+path+`archives/00000005.html" title="02/05/17: #AltWoke 2">link</a>][<a href="`+path+`archives/00000005.html#comments" onmouseover="window.status='Add a comment to this entry';return true" onmouseout="window.status='';return true">No Comments</a>]</div>
</div> <!-- end content -->
<hr>
</div> <!-- end post -->
<br>
}*/

func parse_entries(filename string) (string, string, string, string, string, string, string, []string, string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	metadata := strings.Split(scanner.Text(), "¦")

	postNum := metadata[0]
	name := metadata[1]
	subject := metadata[2]
	epoch, err := strconv.ParseInt(metadata[3], 10, 64)
	if err != nil {
		panic(err)
	}
	datetime := time.Unix(epoch, 0).Format(date_format)
	//datetime := strconv.Itoa(time.Unix(epoch, 0).Year())
	/*archive_name := strconv.Itoa(time.Unix(epoch, 0).Year())  +"/" +
		strconv.Itoa(int(time.Unix(epoch, 0).Month()))*/
	archive_name := strconv.Itoa(time.Unix(epoch, 0).Year())  + "/" +
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

func parse_entries_archive() string {
	c := 0
	filename := "archives/" +strconv.Itoa(c)+".entry"
	toReturn := ""
	_, err := os.Stat(filename);
	for err == nil {
		postNum, _, subject, datetime, _, _, _, _, _ := parse_entries(filename)
		toReturn += `<a href="`+path+`archives/`+postNum+`.html">`+datetime+`: `+subject+`</a><br>`
		c++
		filename = "archives/" +strconv.Itoa(c)+".entry"
		_, err = os.Stat(filename);
	}

	return toReturn;
}

func parse_months_archive() string {
	var extant map[string]bool
	extant = make(map[string]bool)
	c := 0
	filename := "archives/" +strconv.Itoa(c)+".entry"
	toReturn := ""
	_, err := os.Stat(filename);
	for err == nil {
		_, _, _, datetime, _, _, _, _, _ := parse_entries(filename)
		t, errr:= time.Parse(date_format, datetime)
		if errr != nil {
			fmt.Println(errr)
		}
		month := t.Month().String()
		year := strconv.Itoa(t.Year())
		if(!extant[month+"/"+year]) {
			extant[month+"/"+year] = true
			toReturn += `<a href="archives/`+ year +"/"+ month +`.html">`+ month + " " + year +`</a><br>`
		}

		c++
		filename = "archives/" +strconv.Itoa(c)+".entry"
		_, err = os.Stat(filename);
	}

	return toReturn;
}

func parse_comments(c []string) string {
	var toReturn string
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



		toReturn += `<div id="comment`+strconv.Itoa(index)+`"><div id="facesBox`+strconv.Itoa(index)+`" class="facesBox"><div id="picons`+strconv.Itoa(index)+`" class="picons">`+strings.Replace(strings.Trim(fmt.Sprint(search_picons(cmt_email)), "[]"), "> ", ">", -1)+`</div></div><p style="margin-top:0" align="left"> on `+cmt_datetime+`, <a href="`+cmt_hmpg+`" target="_new">`+cmt_name+`</a> [<a href="mailto:`+cmt_email+`" rel="nofollow">e-mail</a>] said
</p><p align="justify">
`+cmt_content+`
</p></div><script>doGravatar("`+cmt_email+`");`
		if cmt_face != "" {
			toReturn += "doFace(\""+cmt_face+"\");"
		}
		if cmt_xface != "" {
			toReturn += "doXFace(\""+cmt_xface+"\");"
		}
		toReturn += "gCount++;</script><hr>"

	}
	return toReturn
}

func search_picons(s string) []string {
	var pBox []string
	if(s=="") {
		pImg := `<img class="face" src="face/picons/misc/MISC/noface/face.gif" title="picon">`
		pBox = append(pBox, pImg)
	} else {
		atSign := strings.Index(s, "@");
		//var mfPiconDatabases = new Array("domains/", "users/", "misc/", "usenix/", "unknown/");
		mfPiconDatabases := [3]string{"domains/", "users/", "usenix/"}
		count := 0
		if (atSign != -1) { // if we have a valid e-mail address..
			host := s[atSign + 1:len(s)]
			user := s[0:atSign]
			host_pieces := strings.Split(host, ".")

			pDef := `<img class="face" src="`+path+`face/picons/unknown/`+host_pieces[len(host_pieces)-1]+`/unknown/face.gif" title="`+host_pieces[len(host_pieces)-1]+`">`
			pBox = append(pBox, pDef)

			for i := range mfPiconDatabases {
				p_path := "face/picons/" + mfPiconDatabases[i]; // they are stored in $PROFILEPATH$/messagefaces/picons/ by default
				if mfPiconDatabases[i] == "misc/" {
					p_path += "MISC/"
				} // special case MISC

				var l = len(host_pieces)-1 // get number of database folders (probably six, but could theoretically change)
				// we will check to see if we have a match at EACH depth, so keep a cloned version w/o the 'unknown/face.gif' portion
				for l >= 0 { // loop through however many pieces we have of the host
					p_path += host_pieces[l]+"/" // add that portion of the host (ex: 'edu' or 'gettysburg' or 'cs')
					clonedLocal := p_path
					if mfPiconDatabases[i] == "users/" {
						p_path += user+"/"
					} else {
						p_path += "unknown/"
					}
					p_path += "face.gif"
					if _, err := os.Stat(p_path); err == nil {
						if(count==0) {
							pBox[0] = `<img class="face" src="`+path+p_path+`"`
							if strings.Contains(p_path,"users") {
								pBox[0] += ` title="`+host_pieces[len(host_pieces)-1]+`">`
							} else {
								pBox[0] += ` title="`+host_pieces[l]+`">`
							}
						} else {
							pImg := `<img class="face" src="`+path+p_path+`"`
							if strings.Contains(p_path, "users") {
								pImg += ` title="`+user+`">`
							} else {
								pImg += ` title="`+host_pieces[l]+`">`
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

func to_markdown(s string) string {
	return strings.Replace(s, "|*|", "<br>", -1)
}

func parse_emoticons(s string) string {
	e_path := "<img src="+path+"emoticons/"
	s = strings.Replace(s,":angry:",e_path+"angry.gif>",-1)
	s = strings.Replace(s,">:(",e_path+"angry.gif>",-1)
	s = strings.Replace(s,":laugh:",e_path+"laugh.gif>",-1)
	s = strings.Replace(s,":DD",e_path+"laugh.gif>",-1)
	s = strings.Replace(s,":yell:",e_path+"yell.gif>",-1)
	s = strings.Replace(s,">:O",e_path+"yell.gif>",-1)
	s = strings.Replace(s,":innocent:",e_path+"innocent.gif>",-1)
	s = strings.Replace(s,"O:)",e_path+"innocent.gif>",-1)
	s = strings.Replace(s,":satisfied:",e_path+"satisfied.gif>",-1)
	s = strings.Replace(s,"/:D",e_path+"satisfied.gif>",-1)
	s = strings.Replace(s,":)",e_path+"smile.gif>",-1)
	s = strings.Replace(s,":O",e_path+"shocked.gif>",-1)
	s = strings.Replace(s,":(",e_path+"sad.gif>",-1)
	s = strings.Replace(s,":D",e_path+"biggrin.gif>",-1)
	s = strings.Replace(s,":P",e_path+"tongue.gif>",-1)
	s = strings.Replace(s,";)",e_path+"wink.gif>",-1)
	s = strings.Replace(s,":blush",e_path+"blush.gif>",-1)
	s = strings.Replace(s,":\")",e_path+"blush.gif>",-1)
	s = strings.Replace(s,":confused:",e_path+"confused.gif>",-1)
	s = strings.Replace(s,":S",e_path+"confused.gif>",-1)
	s = strings.Replace(s,":cool:",e_path+"cool.gif>",-1)
	s = strings.Replace(s,"B)",e_path+"cool.gif>",-1)
	s = strings.Replace(s,":crazy:",e_path+"crazy.gif>",-1)
	s = strings.Replace(s,":cry:",e_path+"cry.gif>",-1)
	s = strings.Replace(s,":~(",e_path+"cry.gif>",-1)
	s = strings.Replace(s,":doze",e_path+"doze.gif>",-1)
	s = strings.Replace(s,":?",e_path+"doze.gif>",-1)
	s = strings.Replace(s,":hehe:",e_path+"hehe.gif>",-1)
	s = strings.Replace(s,"XD",e_path+"hehe.gif>",-1)
	s = strings.Replace(s,":plain:",e_path+"plain.gif>",-1)
	s = strings.Replace(s,":|",e_path+"plain.gif>",-1)
	s = strings.Replace(s,":rolleyes:",e_path+"rolleyes.gif>",-1)
	s = strings.Replace(s,"9_9",e_path+"rolleyes.gif>",-1)
	s = strings.Replace(s,":dizzy:",e_path+"crazy.gif>",-1)
	s = strings.Replace(s,"o_O",e_path+"crazy.gif>",-1)
	s = strings.Replace(s,":money:",e_path+"money.gif>",-1)
	s = strings.Replace(s,":$",e_path+"money.gif>",-1)
	s = strings.Replace(s,":sealed:",e_path+"sealed.gif>",-1)
	s = strings.Replace(s,":X",e_path+"sealed.gif>",-1)
	s = strings.Replace(s,":eek:",e_path+"eek.gif>",-1)
	s = strings.Replace(s,"O_O",e_path+"eek.gif>",-1)
	s = strings.Replace(s,":kiss:",e_path+"kiss.gif>",-1)
	s = strings.Replace(s,":*",e_path+"kiss.gif>",-1)

	return s
}

