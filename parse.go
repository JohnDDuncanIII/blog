package main

import (
	"bufio"
	"fmt"
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

func parse_comments(c []string) string {
	var toReturn string
	for index := range c {
		cmt_splt := strings.Split(c[index],"|")
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

		toReturn += `<div id="comment`+strconv.Itoa(index)+`"><div id="facesBox`+strconv.Itoa(index)+`" class="facesBox"><div id="picons`+strconv.Itoa(index)+`" class="picons">`+strings.Replace(strings.Trim(fmt.Sprint(search_picons(cmt_email)), "[]"), "> ", ">", -1)+`</div></div><p style="margin-top:0" align="left"> on `+cmt_datetime+`, <a href="`+cmt_hmpg+`" target="_new">`+cmt_name+`</a> [<a href="mailto:`+cmt_email+`" rel="nofollow">e-mail</a>] said
</p><p align="justify">
`+cmt_content+`
</p></div><script>doGravatar("`+cmt_email+`");</script><hr>
`
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

			pDef := `<img class="face" src="face/picons/unknown/`+host_pieces[len(host_pieces)-1]+`/unknown/face.gif" title="`+host_pieces[len(host_pieces)-1]+`">`
			pBox = append(pBox, pDef)

			for i := range mfPiconDatabases {
				path := "face/picons/" + mfPiconDatabases[i]; // they are stored in $PROFILEPATH$/messagefaces/picons/ by default
				if mfPiconDatabases[i] == "misc/" {
					path += "MISC/"
				} // special case MISC

				var l = len(host_pieces)-1 // get number of database folders (probably six, but could theoretically change)
				// we will check to see if we have a match at EACH depth, so keep a cloned version w/o the 'unknown/face.gif' portion
				for l >= 0 { // loop through however many pieces we have of the host
					path += host_pieces[l]+"/" // add that portion of the host (ex: 'edu' or 'gettysburg' or 'cs')
					clonedLocal := path
					if mfPiconDatabases[i] == "users/" {
						path += user+"/"
					} else {
						path += "unknown/"
					}
					path += "face.gif"
					if _, err := os.Stat(path); err == nil {
						if(count==0) {
							pBox[0] = `<img class="face" src="`+path+`"`
							if strings.Contains(path,"users") {
								pBox[0] += ` title="`+host_pieces[len(host_pieces)-1]+`">`
							} else {
								pBox[0] += ` title="`+host_pieces[l]+`">`
							}
						} else {
							pImg := `<img class="face" src="`+path+`"`
							if strings.Contains(path, "users") {
								pImg += ` title="`+user+`">`
							} else {
								pImg += ` title="`+host_pieces[l]+`">`
							}
							pBox = append(pBox, pImg)
						}
						count++;
					}
					path = clonedLocal;
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

	file, err := os.Open("entry")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	metadata := strings.Split(scanner.Text(), "|")

	postNum := metadata[0]
	name := metadata[1]
	subject := metadata[2]
	epoch, err := strconv.ParseInt(metadata[3], 10, 64)
	if err != nil {
		panic(err)
	}
	datetime := time.Unix(epoch, 0).Format(date_format)
	//datetime := strconv.Itoa(time.Unix(epoch, 0).Year())

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

	//fmt.Println(metadata,"\n"+ip+"\n"+content+"\n"+more_content)
	//fmt.Println(parse_comments(comments))

	html :=
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
<div class="path"><a href="`+path+`" title="back to frontpage">Home</a> &raquo; <a href="`+path+`archives/" title="weblog archives">Archives</a> &raquo; <a href="`+path+`archives/archive-022017.html" title="archive of February 2017">February 2017</a> &raquo; espy</div>
<div class="direction">
[<a href="http://localhost/~Kelly/DarkMatter/archives/00000001.html">Previous entry: "CujoChat: IRC for Magic Cap!"</a>] [<a href="http://localhost/~Kelly/DarkMatter/archives/00000003.html">Next entry: "Dorian Gray"</a>]
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
<form action="`+path+`cgi-bin/gm-comments.cgi#comments" method="post" name="newcomment">
<table cellpadding="0" cellspacing="2">
<tr>
<td align="center" colspan="2">
<input name="newcommententrynumber" type="hidden" value="2" />
<span style="font-weight:bold;">New Comment</span>
</td>
</tr>
<tr>
<td align="right">Name: </td>
<td><input name="newcommentauthor" size="30" type="text" class="text" /></td>
</tr>
<tr>
<td align="right">E-Mail: </td>
<td><input name="newcommentemail" size="30" type="text" class="text" /></td>
</tr>
<tr>
<td align="right">X-Face: </td>
<td><input name="newcommentxface" size="30" type="text" class="text" /></td>
</tr>
<tr>
<td align="right">Face: </td>
<td><input name="newcommentface" size="30" type="text" class="text" /></td>
</tr>
<tr>
<td align="right">Homepage: </td>
<td><input name="newcommenthomepage" size="30" type="text" class="text" /></td>
</tr>
<tr>
<td>
<table cellpadding="0" cellspacing="0">
<tr>
<td align="center" colspan="3">Smilies:</td>
</tr>
<tr>
<td><img onclick="commentEmoticon(':)')" src="`+path+`emoticons/smile.gif" alt="smile" /></td>
<td><img onclick="commentEmoticon(':O')" src="`+path+`emoticons/shocked.gif" alt="shocked" /></td>
<td><img onclick="commentEmoticon(':(')" src="`+path+`emoticons/sad.gif" alt="sad" /></td>
</tr>
<tr>
<td><img onclick="commentEmoticon(':D')" src="`+path+`emoticons/biggrin.gif" alt="big grin" /></td>
<td><img onclick="commentEmoticon(':P')" src="`+path+`emoticons/tongue.gif" alt="razz" /></td>
<td><img onclick="commentEmoticon(';)')" src="`+path+`emoticons/wink.gif" alt="*wink wink* hey baby" /></td>
</tr>
<tr>
<td><img onclick="commentEmoticon(':angry:')" src="`+path+`emoticons/angry.gif" alt="angry, grr" /></td>
<td><img onclick="commentEmoticon(':blush:')" src="`+path+`emoticons/blush.gif" alt="blush" /></td>
<td><img onclick="commentEmoticon(':confused:')" src="`+path+`emoticons/confused.gif" alt="confused" /></td>
</tr>
<tr>
<td><img onclick="commentEmoticon(':cool:')" src="`+path+`emoticons/cool.gif" alt="cool" /></td>
<td><img onclick="commentEmoticon(':crazy:')" src="`+path+`emoticons/crazy.gif" alt="crazy" /></td>
<td><img onclick="commentEmoticon(':cry:')" src="`+path+`emoticons/cry.gif" alt="cry" /></td>
</tr>
<tr>
<td><img onclick="commentEmoticon(':doze:')" src="`+path+`emoticons/doze.gif" alt="sleepy" /></td>
<td><img onclick="commentEmoticon(':hehe:')" src="`+path+`emoticons/hehe.gif" alt="hehe" /></td>
<td><img onclick="commentEmoticon(':laugh:')" src="`+path+`emoticons/laugh.gif" alt="LOL" /></td>
</tr>
<tr>
<td><img onclick="commentEmoticon(':plain:')" src="`+path+`emoticons/plain.gif" alt="plain jane" /></td>
<td><img onclick="commentEmoticon(':rolleyes:')" src="`+path+`emoticons/rolleyes.gif" alt="rolls eyes" /></td>
<td><img onclick="commentEmoticon(':satisfied:')" src="`+path+`emoticons/satisfied.gif" alt="satisfied" /></td>
</tr>
</table>
</td>
<td>
<textarea cols="25" name="newcommentbody" rows="10" class="text"></textarea>
</td>
</tr>
<tr>
<td>&nbsp;</td>
<td align="center">
<input id="bakecookie" name="bakecookie" type="checkbox" />Save Info?
<label for="bakecookie">
<span></span>
</label>
<br />
<input type="reset" value="Reset" class="button" />
<input name="gmpostpreview" type="submit" value="Preview" class="button" />
<input type="submit" value="Submit" class="button"  onClick="javascript:setGMlocalStorage()" />
</td>
</tr>
</table>
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

<a href="http://greymatterforum.proboards.com/" target="_blank">Greymatter Forums</a></div>
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
<a href="http://greymatterforum.proboards.com/" target="_top"><img src="`+path+`img/dm_1.8.3.gif" alt="Powered By Greymatter" /></a><a href="http://validator.w3.org/check/referer"><img src="`+path+`img/w3c.png" alt="Valid HTML5!"></a>
</div>
</div>
</div>
<script src="`+path+`js/scroll.js"></script>
</body>`
fmt.Println(html)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

