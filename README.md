golang re-implementation of Noah Grey's Greymatter weblogging software

# DONE

* parsing of entries (flat file), with HTML output for individual posts (including links to previous/next posts, [if extant], comments, and post metadata), month archive, archive, and index

* js that adds the ability to hide/show the page header, sidebar, the scroll % of the entry (displayed and updated at the very top of the page), and entry comments.

* support for name, email, homepage, X-Face, and Face in comment submission

* display/decode for [x-face](http://users.cecs.anu.edu.au/~jaa/), [face](http://quimby.gnus.org/circus/face/), [gravatar](http://en.gravatar.com/), and [picons](http://kinzler.com/ftp/faces/picons/) in comment front-end

# TODO:

* administrator front-end

* comment submission

* site search

* separate output for posts made by specific user

* support for separate user repositories through [upspin](https://github.com/upspin/upspin)

* threaded comments

* move gravatar/X-Face parsing to back-end (currently parses every time through js)

* add some sort of Face/X-Face standard for upspin identifiers (maybe unrelated to this project)





