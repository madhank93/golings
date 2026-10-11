// modules5
// The app imports example.com/mathx, a separate module that sits next to it in
// ./mathx with its own go.mod. It is not published anywhere, so go.mod has to
// do two things: require the module, and say where on disk to find it.


module example.com/app

go 1.26.0

require example.com/mathx v0.0.0

replace example.com/mathx => ./mathx
