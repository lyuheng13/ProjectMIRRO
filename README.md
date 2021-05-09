# ProjectMIRRO

## Abstract

This is web project that aims to provide a multi media recommendation system.
From the website, Users are able to find books, music, movie, etc...  
Current the product resources are from other website, but the platform itself will
have the system to track user interests and its own recommendation algorithm.

## APIs

  ### "POST /user/signup"
  #### A sign up handler used to create and register new user to the database. A userid will be returned.  
  
  Req Body:  
    {  
     Email \ string  
     Username \ string  
     Password \ string (more than 8 characters)  
     PasswordConf \ string (Same as Password)  
     Firstname \ string  
     Lastname \ string   
    }  
  
    
  ........................................................................................................................  
  ### "POST /user/login"
  #### Login handler userd to login to the account. A userid will be returned.  
  
  Req Body:  
  {  
    Email \ string  
    Password \ string  
  }  
  
  
  ........................................................................................................................
  ### "GET /user?userid=xxxx"
  #### User Get handler used by the front-end system to get user information.  
  
    
  
  
  ........................................................................................................................
  ### "PATCH /user?userid=xxxx"
  #### User Patch handler used to modify user information in the database.  
  
  Req Body:  
  {  
    Email \ string (Optional)  
    Username \ string (Optional)  
    Password \ string (more than 8 characters) (Optional)  
    Firstname \ string (Optional)  
    Lastname \ string (Optional)  
  }  

  
  
  ........................................................................................................................
    
  ### "DELETE /user?userid=xxxx"
  #### User Delete handler used to delete the user account  
  
    

  ........................................................................................................................
  
  ### "GET /reco?userid=xxxx"
  #### Recommendation handler used to get recommended products.  
  
  Resp (currently):  
  {  
    Youtube \ []string (a list of recommended vedio)  
  }  

  ........................................................................................................................
      
  ### "POST /rate"
  Use: Rate handler used to rate a certain type of products. The recommendation algorithm will update based on user's ratings.  

  Req Body:  
  {  
    Type \ string  
    Like \ Boolean  
    Score \ Integer (From 0 to 10, 10 means perfect)  
    Link \ string (Optional, the product link)  
  }  
  ........................................................................................................................
  
  
## System structure



