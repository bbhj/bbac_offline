ALTER TABLE users change column  open_id openid varchar(255);
ALTER TABLE users change column  union_id unionid varchar(255);



ALTER TABLE images change column  open_id openid varchar(255);
ALTER TABLE images change column  union_id unionid varchar(255);


ALTER TABLE template_form_ids change column  open_id openid varchar(255);
ALTER TABLE template_form_ids change column  form_id formid varchar(255);


ALTER TABLE logins change column  open_id openid varchar(255);
ALTER TABLE logins change column  union_id unionid varchar(255);

ALTER TABLE articles change column  birthed_at birthed_at datetime;
ALTER TABLE articles change column  missed_at missed_at datetime;
