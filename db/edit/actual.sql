ALTER TABLE voices ADD COLUMN eastus BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE voices ADD COLUMN germanywestcentral BOOLEAN NOT NULL DEFAULT false;

UPDATE voices SET eastus = true;