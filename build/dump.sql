PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE voices (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	code_name TEXT NOT NULL UNIQUE,
	is_male BOOLEAN NOT NULL,
	rate INTEGER NOT NULL DEFAULT -1,
	rating INTEGER NOT NULL DEFAULT 0,
	excluded BOOLEAN NOT NULL DEFAULT false,
	comment TEXT NOT NULL DEFAULT ''
	, eastus BOOLEAN NOT NULL DEFAULT false, germanywestcentral BOOLEAN NOT NULL DEFAULT false);
INSERT INTO voices VALUES(1,'Ava Dragon HD Latest','en-US-Ava:DragonHDLatestNeural',0,-1,-2,0,'',1,0);
INSERT INTO voices VALUES(2,'Andrew Dragon HD Latest','en-US-Andrew:DragonHDLatestNeural',1,-1,-5,0,'',1,0);
INSERT INTO voices VALUES(3,'Adam Dragon HD Latest','en-US-Adam:DragonHDLatestNeural',1,-1,-3,0,'',1,0);
INSERT INTO voices VALUES(4,'Alloy Dragon HD Latest','en-US-Alloy:DragonHDLatestNeural',1,-1,-3,0,'',1,0);
INSERT INTO voices VALUES(5,'Aria Dragon HD Latest','en-US-Aria:DragonHDLatestNeural',0,-1,0,0,'',1,0);
INSERT INTO voices VALUES(6,'Bree Dragon HD Latest','en-US-Bree:DragonHDLatestNeural',0,-1,0,0,'',1,0);
INSERT INTO voices VALUES(7,'Brian Dragon HD Latest','en-US-Brian:DragonHDLatestNeural',1,-1,-1,0,'',1,0);
INSERT INTO voices VALUES(8,'Davis Dragon HD Latest','en-US-Davis:DragonHDLatestNeural',1,-1,-1,0,'',1,0);
INSERT INTO voices VALUES(9,'Emma Dragon HD Latest','en-US-Emma:DragonHDLatestNeural',0,-1,-1,0,'',1,0);
INSERT INTO voices VALUES(10,'Emma2 Dragon HD Latest','en-US-Emma2:DragonHDLatestNeural',0,-1,-2,0,'',1,0);
INSERT INTO voices VALUES(11,'Jane Dragon HD Latest','en-US-Jane:DragonHDLatestNeural',0,-1,0,0,'',1,0);
INSERT INTO voices VALUES(12,'Jenny Dragon HD Latest','en-US-Jenny:DragonHDLatestNeural',0,-1,-1,0,'',1,0);
INSERT INTO voices VALUES(13,'Nova Dragon HD Latest','en-US-Nova:DragonHDLatestNeural',0,-1,-1,0,'',1,0);
INSERT INTO voices VALUES(14,'Phoebe Dragon HD Latest','en-US-Phoebe:DragonHDLatestNeural',0,-1,0,0,'',1,0);
INSERT INTO voices VALUES(15,'Serena Dragon HD Latest','en-US-Serena:DragonHDLatestNeural',0,-1,-5,0,'',1,0);
INSERT INTO voices VALUES(16,'Steffan Dragon HD Latest','en-US-Steffan:DragonHDLatestNeural',1,-1,-2,0,'',1,0);
INSERT INTO voices VALUES(17,'Andrew DragonHD Omni Latest','en-US-Andrew:DragonHDOmniLatestNeural',1,-1,-2,0,'',1,0);
INSERT INTO voices VALUES(18,'Caleb DragonHD Omni Latest','en-US-Caleb:DragonHDOmniLatestNeural',1,-1,-1,0,'',1,0);
INSERT INTO voices VALUES(19,'Dana DragonHD Omni Latest','en-US-Dana:DragonHDOmniLatestNeural',0,-1,0,0,'',1,0);
INSERT INTO voices VALUES(20,'Lewis DragonHD Omni Latest','en-US-Lewis:DragonHDOmniLatestNeural',1,-1,-2,0,'',1,0);
INSERT INTO voices VALUES(21,'Phoebe DragonHD Omni Latest','en-US-Phoebe:DragonHDOmniLatestNeural',0,-1,0,0,'',1,0);
INSERT INTO voices VALUES(22,'Ava','en-US-AvaNeural',0,-1,0,0,'',1,1);
INSERT INTO voices VALUES(23,'Andrew','en-US-AndrewNeural',1,-1,-1,0,'',1,1);
INSERT INTO voices VALUES(24,'Emma','en-US-EmmaNeural',0,-1,-1,0,'',1,1);
INSERT INTO voices VALUES(25,'Brian','en-US-BrianNeural',1,-1,-1,0,'',1,1);
INSERT INTO voices VALUES(26,'Jenny','en-US-JennyNeural',0,152,0,0,'',1,1);
INSERT INTO voices VALUES(27,'Guy','en-US-GuyNeural',1,215,0,0,'',1,1);
INSERT INTO voices VALUES(28,'Aria','en-US-AriaNeural',0,150,0,0,'',1,1);
INSERT INTO voices VALUES(29,'Davis','en-US-DavisNeural',1,154,-1,0,'',1,1);
INSERT INTO voices VALUES(30,'Jane','en-US-JaneNeural',0,154,-1,0,'',1,1);
INSERT INTO voices VALUES(31,'Jason','en-US-JasonNeural',1,156,0,0,'',1,1);
INSERT INTO voices VALUES(32,'Kai','en-US-KaiNeural',1,-1,-2,0,'',1,1);
INSERT INTO voices VALUES(33,'Luna','en-US-LunaNeural',0,-1,0,0,'',1,1);
INSERT INTO voices VALUES(34,'Sara','en-US-SaraNeural',0,157,0,0,'',1,1);
INSERT INTO voices VALUES(35,'Tony','en-US-TonyNeural',1,156,0,0,'',1,1);
INSERT INTO voices VALUES(36,'Nancy','en-US-NancyNeural',0,149,0,0,'',1,1);
INSERT INTO voices VALUES(37,'Amber','en-US-AmberNeural',0,152,0,0,'',1,1);
INSERT INTO voices VALUES(38,'Ana','en-US-AnaNeural',0,135,0,0,'',1,1);
INSERT INTO voices VALUES(39,'Ashley','en-US-AshleyNeural',0,149,0,0,'',1,1);
INSERT INTO voices VALUES(40,'Brandon','en-US-BrandonNeural',1,156,0,0,'',1,1);
INSERT INTO voices VALUES(41,'Christopher','en-US-ChristopherNeural',1,149,0,0,'',1,1);
INSERT INTO voices VALUES(42,'Cora','en-US-CoraNeural',0,146,0,0,'',1,1);
INSERT INTO voices VALUES(43,'Elizabeth','en-US-ElizabethNeural',0,152,0,0,'',1,1);
INSERT INTO voices VALUES(44,'Eric','en-US-EricNeural',1,147,0,0,'',1,1);
INSERT INTO voices VALUES(45,'Jacob','en-US-JacobNeural',1,154,0,0,'',1,1);
INSERT INTO voices VALUES(46,'en-US-Jimmie:DragonHDFlashLatestNeural','en-US-Jimmie:DragonHDFlashLatestNeural',1,-1,-6,0,'',1,0);
INSERT INTO voices VALUES(47,'Michelle','en-US-MichelleNeural',0,154,0,0,'',1,1);
INSERT INTO voices VALUES(48,'Monica','en-US-MonicaNeural',0,145,0,0,'',1,1);
INSERT INTO voices VALUES(49,'Roger','en-US-RogerNeural',1,-1,0,0,'',1,1);
INSERT INTO voices VALUES(50,'Steffan','en-US-SteffanNeural',1,154,0,0,'',1,1);
INSERT INTO voices VALUES(51,'en-US-Tiana:DragonHDFlashLatestNeural','en-US-Tiana:DragonHDFlashLatestNeural',0,-1,0,0,'',1,0);
INSERT INTO voices VALUES(52,'en-US-Tyler:DragonHDFlashLatestNeural','en-US-Tyler:DragonHDFlashLatestNeural',1,-1,-2,0,'',1,0);
INSERT INTO voices VALUES(53,'AIGenerate1','en-US-AIGenerate1Neural',1,135,0,0,'',1,0);
INSERT INTO voices VALUES(54,'AIGenerate2','en-US-AIGenerate2Neural',0,140,-1,0,'',1,0);
INSERT INTO voices VALUES(55,'Andrew2 Dragon HD Latest','en-US-Andrew2:DragonHDLatestNeural',1,-1,-2,0,'',1,0);
INSERT INTO voices VALUES(56,'Andrew3 Dragon HD Latest','en-US-Andrew3:DragonHDLatestNeural',1,-1,-1,0,'',1,0);
INSERT INTO voices VALUES(57,'Ava DragonHD Omni Latest','en-us-ava:DragonHDOmniLatestNeural',0,-1,-1,0,'',1,0);
INSERT INTO voices VALUES(58,'Ava3 Dragon HD Latest','en-US-Ava3:DragonHDLatestNeural',0,-1,-1,0,'',1,0);
CREATE TABLE records (
		expression TEXT PRIMARY KEY UNIQUE,
		v1 TEXT NOT NULL DEFAULT '',
		v2 TEXT NOT NULL DEFAULT '',
		v3 TEXT NOT NULL DEFAULT '',
		v4 TEXT NOT NULL DEFAULT '',
		v5 TEXT NOT NULL DEFAULT '',
		v6 TEXT NOT NULL DEFAULT ''
	) WITHOUT ROWID;
INSERT INTO records VALUES('August','','','en-US-Caleb:DragonHDOmniLatestNeural','en-US-Nova:DragonHDLatestNeural','en-US-AndrewNeural','');
INSERT INTO records VALUES('February','en-US-Adam:DragonHDLatestNeural','en-US-EmmaNeural','en-US-AIGenerate1Neural','en-US-Serena:DragonHDLatestNeural','en-US-Steffan:DragonHDLatestNeural','en-US-Aria:DragonHDLatestNeural');
INSERT INTO records VALUES('July','','en-US-SaraNeural','en-US-GuyNeural','en-US-AmberNeural','en-US-Davis:DragonHDLatestNeural','');
INSERT INTO records VALUES('March','en-US-AIGenerate1Neural','en-US-NancyNeural','en-US-Steffan:DragonHDLatestNeural','en-US-Jane:DragonHDLatestNeural','','');
INSERT INTO records VALUES('ado','en-US-SteffanNeural','en-US-ElizabethNeural','','','','');
INSERT INTO records VALUES('advisor','en-US-Andrew:DragonHDOmniLatestNeural','en-US-MonicaNeural','','','','');
INSERT INTO records VALUES('aforementioned','en-US-Lewis:DragonHDOmniLatestNeural','en-US-Ava3:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('aftermath','en-US-Lewis:DragonHDOmniLatestNeural','en-US-LunaNeural','','','','');
INSERT INTO records VALUES('ain''t','','','en-US-Tyler:DragonHDFlashLatestNeural','','','');
INSERT INTO records VALUES('ascribe','en-US-BrianNeural','en-US-Nova:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('at a loss for words','en-US-SteffanNeural','en-US-Emma:DragonHDLatestNeural','','','en-US-Tyler:DragonHDFlashLatestNeural','en-US-MonicaNeural');
INSERT INTO records VALUES('avail','en-US-Davis:DragonHDLatestNeural','en-US-AnaNeural','','','','');
INSERT INTO records VALUES('brunt','en-US-ChristopherNeural','en-US-AshleyNeural','','','','');
INSERT INTO records VALUES('busted','en-US-Jimmie:DragonHDFlashLatestNeural','en-US-Dana:DragonHDOmniLatestNeural','','','','');
INSERT INTO records VALUES('but then','en-US-TonyNeural','en-US-CoraNeural','en-US-Andrew3:DragonHDLatestNeural','en-US-JennyNeural','en-US-BrandonNeural','en-US-EmmaNeural');
INSERT INTO records VALUES('carry out','en-US-Lewis:DragonHDOmniLatestNeural','en-US-Emma:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('conspectus','en-US-SteffanNeural','en-US-Ava:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('cuz','en-US-TonyNeural','en-US-AmberNeural','en-US-Brian:DragonHDLatestNeural','','en-US-DavisNeural','en-US-Emma:DragonHDLatestNeural');
INSERT INTO records VALUES('endowed','en-US-RogerNeural','en-US-Ava:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('enrapture','','en-US-AmberNeural','','','','');
INSERT INTO records VALUES('fed','en-US-Andrew:DragonHDOmniLatestNeural','en-US-Ava:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('flowchart','en-US-KaiNeural','en-US-NancyNeural','','','','');
INSERT INTO records VALUES('for good','en-US-Lewis:DragonHDOmniLatestNeural','en-US-SaraNeural','','','','');
INSERT INTO records VALUES('forcibly','en-US-Andrew2:DragonHDLatestNeural','en-US-Tiana:DragonHDFlashLatestNeural','','','','');
INSERT INTO records VALUES('gist','en-US-Andrew:DragonHDLatestNeural','en-US-Nova:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('give away','en-US-JasonNeural','en-US-EmmaNeural','','','','');
INSERT INTO records VALUES('go crazy','en-US-Davis:DragonHDLatestNeural','en-US-JennyNeural','en-US-KaiNeural','en-US-AvaNeural','','en-US-Aria:DragonHDLatestNeural');
INSERT INTO records VALUES('help out','','en-US-MonicaNeural','','','','');
INSERT INTO records VALUES('homeroom teacher','','','','','en-US-Lewis:DragonHDOmniLatestNeural','');
INSERT INTO records VALUES('in effect','en-US-ChristopherNeural','en-US-ElizabethNeural','','','','');
INSERT INTO records VALUES('in good faith','en-US-Andrew2:DragonHDLatestNeural','en-US-AriaNeural','en-US-EricNeural','en-US-EmmaNeural','en-US-KaiNeural','en-US-Phoebe:DragonHDLatestNeural');
INSERT INTO records VALUES('in return','en-US-BrandonNeural','en-US-AnaNeural','','','','');
INSERT INTO records VALUES('indices','en-US-KaiNeural','en-US-Dana:DragonHDOmniLatestNeural','en-US-TonyNeural','en-US-Jenny:DragonHDLatestNeural','en-US-Alloy:DragonHDLatestNeural','en-US-Bree:DragonHDLatestNeural');
INSERT INTO records VALUES('ingratiate','en-US-SteffanNeural','en-US-AvaNeural','','','','');
INSERT INTO records VALUES('jiffy','en-US-RogerNeural','en-US-JaneNeural','','','','');
INSERT INTO records VALUES('judge from','en-US-Lewis:DragonHDOmniLatestNeural','en-US-Tiana:DragonHDFlashLatestNeural','en-US-BrianNeural','en-US-AnaNeural','en-US-AndrewNeural','en-US-Aria:DragonHDLatestNeural');
INSERT INTO records VALUES('knotted','en-US-Jimmie:DragonHDFlashLatestNeural','en-US-JaneNeural','','','','');
INSERT INTO records VALUES('last resort','en-US-Jimmie:DragonHDFlashLatestNeural','en-US-AriaNeural','','','','');
INSERT INTO records VALUES('lest','','','en-US-Andrew:DragonHDOmniLatestNeural','en-US-Serena:DragonHDLatestNeural','','');
INSERT INTO records VALUES('live out','en-US-JacobNeural','en-US-AvaNeural','','','','');
INSERT INTO records VALUES('lovey-dovey','','','en-US-JacobNeural','en-US-CoraNeural','en-US-GuyNeural','en-US-LunaNeural');
INSERT INTO records VALUES('low tide','en-US-Davis:DragonHDLatestNeural','en-US-AvaNeural','','','','');
INSERT INTO records VALUES('make-believe','en-US-EricNeural','en-US-EmmaNeural','','','','');
INSERT INTO records VALUES('make-up','en-US-EricNeural','en-US-EmmaNeural','','','','');
INSERT INTO records VALUES('obstinacy','en-US-JasonNeural','en-US-SaraNeural','','','','');
INSERT INTO records VALUES('outset','en-US-TonyNeural','en-US-AriaNeural','','','','');
INSERT INTO records VALUES('outstretch','en-US-Andrew3:DragonHDLatestNeural','en-US-ElizabethNeural','en-US-Brian:DragonHDLatestNeural','','en-US-Davis:DragonHDLatestNeural','en-US-Bree:DragonHDLatestNeural');
INSERT INTO records VALUES('overridden','en-US-KaiNeural','en-US-JaneNeural','','','en-US-Tyler:DragonHDFlashLatestNeural','en-US-Aria:DragonHDLatestNeural');
INSERT INTO records VALUES('overrode','','','en-US-Caleb:DragonHDOmniLatestNeural','en-US-JennyNeural','en-US-Andrew3:DragonHDLatestNeural','en-US-NancyNeural');
INSERT INTO records VALUES('perplexed','en-US-TonyNeural','en-US-AvaNeural','','','','');
INSERT INTO records VALUES('play along','en-US-Jimmie:DragonHDFlashLatestNeural','en-US-AriaNeural','','','','');
INSERT INTO records VALUES('privy','en-US-JasonNeural','','','','','');
INSERT INTO records VALUES('prolly','en-US-EricNeural','en-US-JennyNeural','en-US-TonyNeural','en-US-Serena:DragonHDLatestNeural','en-US-Andrew:DragonHDOmniLatestNeural','en-US-AshleyNeural');
INSERT INTO records VALUES('prying','en-US-BrianNeural','en-US-MonicaNeural','','','','');
INSERT INTO records VALUES('pull off','en-US-Brian:DragonHDLatestNeural','en-US-Aria:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('reenact','en-US-SteffanNeural','en-US-MichelleNeural','','','','');
INSERT INTO records VALUES('run out','en-US-Andrew:DragonHDLatestNeural','en-US-AIGenerate2Neural','','','','');
INSERT INTO records VALUES('scrambled eggs','en-US-GuyNeural','en-US-Tiana:DragonHDFlashLatestNeural','','','','');
INSERT INTO records VALUES('shed','en-US-Steffan:DragonHDLatestNeural','en-US-AnaNeural','','','','');
INSERT INTO records VALUES('shock wave','en-US-Steffan:DragonHDLatestNeural','en-US-MonicaNeural','','','','');
INSERT INTO records VALUES('sigil','en-US-Lewis:DragonHDOmniLatestNeural','en-US-ElizabethNeural','','','en-US-Andrew:DragonHDLatestNeural','en-US-Nova:DragonHDLatestNeural');
INSERT INTO records VALUES('simultaneously','en-US-Andrew3:DragonHDLatestNeural','en-US-MonicaNeural','','','','');
INSERT INTO records VALUES('smithereens','en-US-GuyNeural','en-US-JennyNeural','','','','');
INSERT INTO records VALUES('sprang','','en-US-AriaNeural','en-US-AIGenerate1Neural','en-US-CoraNeural','en-US-JacobNeural','en-US-Phoebe:DragonHDOmniLatestNeural');
INSERT INTO records VALUES('sprung','','','','en-US-Dana:DragonHDOmniLatestNeural','en-US-BrandonNeural','en-US-Phoebe:DragonHDOmniLatestNeural');
INSERT INTO records VALUES('stand one''s ground','','','en-US-JacobNeural','en-US-Jane:DragonHDLatestNeural','en-US-SteffanNeural','');
INSERT INTO records VALUES('step in','en-US-RogerNeural','en-US-Bree:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('stick out like a sore thumb','en-US-Lewis:DragonHDOmniLatestNeural','en-US-AriaNeural','','','','');
INSERT INTO records VALUES('stolen','en-US-EricNeural','','','en-US-AshleyNeural','en-US-Caleb:DragonHDOmniLatestNeural','en-US-LunaNeural');
INSERT INTO records VALUES('stomach ','en-US-BrandonNeural','en-US-Emma2:DragonHDLatestNeural','en-US-AIGenerate1Neural','en-US-Aria:DragonHDLatestNeural','en-US-RogerNeural','en-US-AnaNeural');
INSERT INTO records VALUES('swiveling','en-US-Caleb:DragonHDOmniLatestNeural','en-US-JennyNeural','en-US-Lewis:DragonHDOmniLatestNeural','en-US-Aria:DragonHDLatestNeural','en-US-BrianNeural','en-US-Phoebe:DragonHDOmniLatestNeural');
INSERT INTO records VALUES('tad','en-US-Lewis:DragonHDOmniLatestNeural','en-US-Aria:DragonHDLatestNeural','','','','');
INSERT INTO records VALUES('thrall','en-US-Davis:DragonHDLatestNeural','en-US-AriaNeural','','','','');
INSERT INTO records VALUES('to be pretty well off','en-US-Andrew3:DragonHDLatestNeural','en-US-ElizabethNeural','en-US-Andrew2:DragonHDLatestNeural','en-US-Emma2:DragonHDLatestNeural','en-US-Alloy:DragonHDLatestNeural','en-US-AshleyNeural');
INSERT INTO records VALUES('to get used','en-US-EricNeural','en-US-Aria:DragonHDLatestNeural','en-US-BrianNeural','en-US-Tiana:DragonHDFlashLatestNeural','en-US-Lewis:DragonHDOmniLatestNeural','en-US-Dana:DragonHDOmniLatestNeural');
INSERT INTO records VALUES('to go all out','en-US-GuyNeural','en-US-SaraNeural','en-US-Andrew:DragonHDOmniLatestNeural','en-US-Phoebe:DragonHDOmniLatestNeural','en-US-Andrew3:DragonHDLatestNeural','en-US-Dana:DragonHDOmniLatestNeural');
INSERT INTO records VALUES('to make up one''s mind','en-US-EricNeural','','en-US-Davis:DragonHDLatestNeural','','en-US-Lewis:DragonHDOmniLatestNeural','en-US-Nova:DragonHDLatestNeural');
INSERT INTO records VALUES('wed','en-US-SteffanNeural','en-US-AvaNeural','','','','');
INSERT INTO records VALUES('weightlessness','en-US-DavisNeural','en-US-AIGenerate2Neural','','','','');
INSERT INTO records VALUES('wept','','','','','','en-US-Jenny:DragonHDLatestNeural');
PRAGMA writable_schema=ON;
CREATE TABLE IF NOT EXISTS sqlite_sequence(name,seq);
DELETE FROM sqlite_sequence;
INSERT INTO sqlite_sequence VALUES('voices',86);
PRAGMA writable_schema=OFF;
COMMIT;
