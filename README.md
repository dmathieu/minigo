# Minigo [![Build Status](https://travis-ci.org/dmathieu/minigo.png)](https://travis-ci.org/dmathieu/minigo)

This is an experiment at writing some Go.  
It's a very basic url shortener which you can deploy to heroku.

## Usage

Go to the app's URI.  
Enter the url you wish to minify in the form field.
You will then see the slug generated for your URI.

Go to your app's URI `/yourslug`. And you will be redirected to the URI you specified

See, for example `http://minigo.herokuapp.com/snqDQ3q8TSIa5sWLTs9lfw` ([link](http://minigo.herokuapp.com/snqDQ3q8TSIa5sWLTs9lfw)), which redirects to dmathieu.com

## Perks

* There is no unicity validation for now. In case a URL gets the same slug as an other one (that's highly improbable though), you might not get redirected to the right place
* This is just an experiment. You shouldn't use it in production. Feel free to hack on it and submit pull requests though.

## Maintainer

* Damien MATHIEU ([github/dmathieu](http://github.com/dmathieu), [dmathieu.com](http://dmathieu.com))

## License
MIT License. Copyright 2013 Damien MATHIEU
