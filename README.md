golang re-implementation of Noah Grey's Greymatter weblogging software

# DONE

* parsing of entries (flat file), with HTML output for individual posts (including links to previous/next posts, [if extant], comments, and post metadata), month archive, archive, and index

* js that adds the ability to hide/show the page header, sidebar, the scroll % of the entry (displayed and updated at the very top of the page), and entry comments.

* support for name, email, homepage, X-Face, and Face in comment submission

* display/decode for [x-face](http://users.cecs.anu.edu.au/~jaa/), [face](http://quimby.gnus.org/circus/face/), [gravatar](http://en.gravatar.com/), and [picons](http://kinzler.com/ftp/faces/picons/) in comment front-end

# TODO:

* administrator front-end (probably clone gm.cgi layout from Greymatter)

* comment submission (see comment.go)

* site search (see search.go)

* separate output for posts made by specific user (allow multiple users to run off of one instance; john.domain.com)

* support for user repositories through [upspin](https://github.com/upspin/upspin)

* threaded comments (looking at a clean way of implementing this w/ the flat-file comment structure)

* RSS/atom (1.0/2.0/0.3) syndication

* move gravatar/X-Face parsing to back-end (currently parses every time through js)

* add some sort of Face/X-Face standard for upspin identifiers (maybe unrelated to this project)

* templating (add [Brad Fitzpatrick's](http://brad.livejournal.com/) default livejournal template; support custom HTML/CSS/JS (to a degree)

* image upload for posts

* tags/categories

* trackbacks/linkbacks

* geolocation

* calendar (for archive and sidebar)

* OpenID

* scheduled posting

# INSPIRATION

[Noah Grey's](http://noahgrey.com/) [Greymatter](https://github.com/JohnDDuncanIII/greymatter/), [Brad Fitzpatrick's](http://bradfitz.com/) [LiveJournal](https://github.com/apparentlymart/livejournal), and [SixApart's](https://www.sixapart.jp/) [MovableType](https://github.com/movabletype/movabletype)