package com.tfm.data.login;

import javax.xml.bind.annotation.XmlRootElement;

import com.tfm.data.federacion.Federacion;
import com.tfm.data.propietario.Propietario;
import com.tfm.data.veterinario.Veterinario;

@XmlRootElement
public class Login {

	private int id;
	private String nickname;
	private String tipo;
	private String email;
	private String password;
	private int telefono;
	private Propietario propietario;
	private Veterinario veterinario;
	private Federacion federacion;

	public Login(int id, String nickname, String tipo, String email, String password, int telefono,
			Propietario propietario, Veterinario veterinario, Federacion federacion) {
		this.id = id;
		this.nickname = nickname;
		this.tipo = tipo;
		this.email = email;
		this.password = password;
		this.telefono = telefono;
		this.propietario = propietario;
		this.veterinario = veterinario;
		this.federacion = federacion;
	}

	public Login(int id, String tipo) {
		this.id = id;
		this.tipo = tipo;
	}

	public Login() {

	}

	public Veterinario getVeterinario() {
		return veterinario;
	}

	public void setVeterinario(Veterinario veterinario) {
		this.veterinario = veterinario;
	}

	public Federacion getFederacion() {
		return federacion;
	}

	public void setFederacion(Federacion federacion) {
		this.federacion = federacion;
	}

	public Propietario getPropietario() {
		return propietario;
	}

	public void setPropietario(Propietario propietario) {
		this.propietario = propietario;
	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public String getNickname() {
		return nickname;
	}

	public void setNickname(String nickname) {
		this.nickname = nickname;
	}

	public String getTipo() {
		return tipo;
	}

	public void setTipo(String tipo) {
		this.tipo = tipo;
	}

	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	public int getTelefono() {
		return telefono;
	}

	public void setTelefono(int telefono) {
		this.telefono = telefono;
	}

}
