<form action="/auth/login" method="post">
    <div class="form-group"></div>
        {{.Form | renderform}}
    </div>
  <button type="submit" class="btn btn-primary">Submit</button>
</form>