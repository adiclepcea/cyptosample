package org.clepcea.secured.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SecuredController {
    @GetMapping({"","/"})
    public String verySecureInfo(){
        return "You are brave!";
    }
}
