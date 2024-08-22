CREATE  TABLE dialogo (
                          id                   uuid  NOT NULL  ,
                          id_usuario           uuid  NOT NULL  ,
                          criado               timestamptz    ,
                          CONSTRAINT pk_dialogo PRIMARY KEY ( id )
);

COMMENT ON TABLE dialogo IS 'Tabela que armazena as informações sobre pesquisas realizadas na Gio';

COMMENT ON COLUMN dialogo.id IS 'Identificador do registro';

COMMENT ON COLUMN dialogo.id_usuario IS 'Identificador de relacionamento';


COMMENT ON COLUMN dialogo.criado IS 'Data de criação do registro';

CREATE  TABLE dialogo_detalhe (
                                  id                   uuid  NOT NULL  ,
                                  id_dialogo           uuid  NOT NULL  ,
                                  pergunta             text  NOT NULL  ,
                                  resposta             text  NOT NULL  ,
                                  criado               timestamptz    ,
                                  CONSTRAINT pk_dialogo_detalhe PRIMARY KEY ( id )
);

CREATE INDEX idx_dialogo_detalhe_01 ON dialogo_detalhe  ( id_dialogo );

ALTER TABLE dialogo_detalhe ADD CONSTRAINT fk_dialogo_detalhe_01 FOREIGN KEY ( id_dialogo ) REFERENCES dialogo( id );

COMMENT ON TABLE dialogo_detalhe IS 'Tabela que armazena as informações sobre os diálogos realizadas na Gio';

COMMENT ON COLUMN dialogo_detalhe.id IS 'Identificador do registro';

COMMENT ON COLUMN dialogo_detalhe.id_dialogo IS 'Identificador de relacionamento';

COMMENT ON COLUMN dialogo_detalhe.pergunta IS 'pergunta realizada para a Gio';

COMMENT ON COLUMN dialogo_detalhe.resposta IS 'Resposta recebida para a pergunta realizada';

COMMENT ON COLUMN dialogo_detalhe.insight IS 'Insight referente a pergunta realizada';

COMMENT ON COLUMN dialogo_detalhe.token IS 'Quantidade de tokens enviados na pergunta';

COMMENT ON COLUMN dialogo_detalhe.criado IS 'Data de criação do registro';