package com.tfm.security;

import java.util.Calendar;
import java.util.Date;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

public class Token {
	
	public static String issueToken(String login) {
		// Calculamos la fecha de expiración del token
		Date issueDate = new Date();
		Calendar calendar = Calendar.getInstance();
		calendar.setTime(issueDate);
		calendar.add(Calendar.MINUTE, 60);
		Date expireDate = calendar.getTime();

		// Creamos el token
		return Jwts.builder().setSubject(login).setIssuer("https://www.kupiri.es").setIssuedAt(issueDate)
				.setExpiration(expireDate).signWith(SignatureAlgorithm.HS512, RestSecurityFilter.KEY).compact();
		
	}

}
