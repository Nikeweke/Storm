// Cron Format - http://www.nncron.ru/help/RU/working/cron-format.htm
// Cron online maker - https://crontab.guru
// example
// '* * * * * *'    - каждую секунду (6 звезд)
// '*/30 * * * * *' - через каждые 30 сек запускать работу (6 звезд)
// '*/1 * * * *'    - через каждую мин. (5 звезд)
package config

import (
	"github.com/robfig/cron"
	"fmt"
	"../app/helpers"
	color "github.com/fatih/color"
)

func Jobs() {
	if helpers.GetSettings("jobs").(bool) == false {
		return
	}

	color.New(color.FgCyan).Add(color.Bold).Print("Jobs")
	color.New(color.FgGreen).Add(color.Bold).Print(" are enabled \n")

  // Jobs 
	c := cron.New()

	// Every second
	c.AddFunc("@every 1s", func() { fmt.Println("Every second") }) 
	// c.AddFunc("* * * * * *", func() { fmt.Println("Every second") })  

	// Every 10 second
	// c.AddFunc("@every 10s", func() { fmt.Println("Every 10 second") }) 
	// c.AddFunc("*/10 * * * * *", func() { fmt.Println("Every 10 second") })  

	// Every 1 minute 
	// c.AddFunc("@every 1m", func() { fmt.Println("Every 1 minute") }) 
	// c.AddFunc("*/60 * * * * *", func() { fmt.Println("Every 1 minute") }) 

	// Every hour 
	// c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	// c.AddFunc("@every 1h",      func() { fmt.Println("Every hour") })
	// c.AddFunc("0 */1 * * *", func() { fmt.Println("Every hour") }) 

	c.Start()
	// c.Stop()
}
