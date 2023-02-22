PROJECT = "gpiodemo"
VERSION = "1.0.0"

--[[
    Sonde Sirene
        surveille le topic alarme_armee
        surveille le topic alarme
		Si les deux topics sont à 1 alors enclancher la sirenne
]]

local sys = require "sys"
require("sysplus")


log.info("main", "sonde_sirene")

if wdt then
    -- Watchdog 
    wdt.init(9000)--初始化watchdog 9s
    sys.timerLoopStart(wdt.feed, 3000)--3s
end
	
local device_id     = "porte-entree"    --Passez à votre propre appareil id
--local device_id     = "votre_id"    --Passez à votre propre appareil id
local device_secret = "votre_cle"    --Changer pour votre propre clé d'appareil

local mqttc = nil

-- gpio12/gpio13
local LEDA= gpio.setup(12, 0, gpio.PULLUP) -- gpio.setup(pin, mode, pull, irq)
local LEDB= gpio.setup(13, 0, gpio.PULLUP)

local table = {"/alarme_armee",9,"/alarme",9}

sys.subscribe("IP_READY", function()
    log.info("mobile", "IP_READY")
	-- LEDB(1) retirer le 22/02/2023
end)

sys.subscribe("WLAN_READY", function ()
    if wlan.ready() == 0 then
            log.info("network", "tsend complete, sleep 5s")
			--LEDB(1)
        else
            log.warn("main", "wlan is not ready yet")
        	--LEDB(0)			
        end	
end)

sys.subscribe("NTP_UPDATE", function()
    log.info("date", "time_ici", os.date())
end)

sys.subscribe("NTP_SYNC_DONE", function()
        log.info("ntp", "done")
        log.info("la date est synchro", os.date())
    end
)

sys.taskInit(function()
    log.info("wlan", "Lancement initialisation Wifi :", wlan.init())
    wlan.setMode(wlan.STATION)
    --wlan.connect("votre_wifi", "votre_id_wifi", 1)
	wlan.connect("votre_wifi", "votre_id_wifi", 1)
    local result, data = sys.waitUntil("IP_READY")
    log.info("wlan", "IP_READY", result, data)
    
    local client_id,user_name,password = iotauth.iotda(device_id,device_secret)
    log.info("iotda",client_id,user_name,password)
	
    mqttc = mqtt.create(nil,"192.168.0.140", 1883)
	--mqttc = mqtt.create(nil,"192.168.1.5", 1883)
    mqttc:auth(client_id,user_name,password)
    mqttc:keepalive(30) --  keepaliv 240s
    mqttc:autoreconn(true, 3000) -- autoreconnexion tous les 3000

    mqttc:on(function(mqtt_client, event, data, payload)
        -- 用户自定义代码
        log.info("mqtt", "event", event, mqtt_client, data, payload)
        if event == "conack" then
            sys.publish("mqtt_conack")
			mqtt_client:subscribe({["/alarme"]=1,["/alarme_armee"]=2})
        elseif event == "recv" then
            log.info("mqtt", "receive", "topic", data, "payload", payload)
			-- Mettre une variable locale representant le topic
			--local topic_recue = {data,payload}
			--Test du topic recu data 
			if data == "/alarme_armee" then
				table[2] = payload
				log.info("Je suis passe par if data /alarme_armee")
            elseif data == "/alarme" then
				table[4] = payload
				log.info("Je suis passe par if data /alarme")
			elseif data ~= "alarme" then
				log.info("Je suis passe par if data diferent de /alarme")
			end
			log.info("table variable globale :", table[1],table[2],table[3],table[4])
        elseif event == "sent" then
            log.info("mqtt", "sent", "pkgid", data)
			
        end
    end)
    mqttc:connect()
    sys.wait(10000)
    --mqttc:subscribe("/alarme/")
	mqttc:subscribe({["/alarme"]=1,["/alarme_armee"]=2})
	sys.waitUntil("mqtt_conack")
    while true do
        -- mqttc自动处理重连
        local ret, topic, data, qos = sys.waitUntil("mqtt_pub", 30000)
        if ret then
            if topic == "close" then break end
            mqttc:publish(topic, data, qos)
        end
    end
    mqttc:close()
    mqttc = nil
end)

sys.taskInit(function()
--    -- publie 1 dans le topic sirene
--	local topic = "/sirene/"
--	local payload = "1"
--	local qos = 1
    while true do
        sys.wait(5000)
        --if mqttc:ready() then
        --    local pkgid = mqttc:publish(topic, payload, qos)
		--	end
		log.info("Tache pour controler la table etposition le GPIO ou la led")
		if table[2] == "1" and table[4]== "1" then
     		LEDB(1)
		else
		    log.info("Tache  variable globale :", table[1],table[2],table[3],table[4])
			LEDB(0)
		end
    end
end) 

--  ---------------------------------------------
-- 
sys.run()
-- sys.run()
