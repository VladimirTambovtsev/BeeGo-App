{{template "_navigation.tpl"}}

<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <h1 class="mt-4 mb-4">Sign In</h1>
            <form method="POST" action="">
                <div class="form-group">
                    <label for="exampleInputEmail1">Email address</label>
                    <input name="email" type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email">
                </div>
                <div class="form-group">
                    <label for="exampleInputPassword1">Password</label>
                    <input name="password" type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
                </div>
                <button type="submit" class="btn btn-primary mt-2 mb-2">Submit</button>
                <div class="form-group">
                    <small>
                        <a href="/account/create">Signup</a>
                    </small>
                </div>
            </form>
        </div>
    </div>
</div>
