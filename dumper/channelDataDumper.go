package dumper

import (
	"database/sql"
	"fmt"

	"github.com/moio/mgr-dump/schemareader"
)

func readTableNames() []string {
	return []string{
		// dictionaries
		"rhnproductname",
		"rhnchannelproduct",
		"rhnarchtype",
		"rhnchecksumtype",
		"rhnpackagearch",
		"web_customer",
		"rhnchannelarch",
		"rhnerrataseverity", // this table is static (even the id's). Should we copy it?
		//8
		// data to transfer
		"rhnchannel",
		"rhnchannelfamily",
		"rhnchannelfamilymembers",
		"rhnerrata",
		"rhnchannelerrata",
		//13
		"rhnpackagename",  // done
		"rhnpackagegroup", // done
		"rhnsourcerpm",    // done
		"rhnpackageevr",   // done
		"rhnchecksum",     // done
		//18
		"rhnpackage",
		"rhnchannelpackage",
		"rhnerratapackage",
		//21
		"rhnpackageprovider", // catalog
		"rhnpackagekeytype",  // catalog
		"rhnpackagekey",      // catalog
		"rhnpackagekeyassociation",
		//25
		"rhnerratabuglist",

		"rhnpackagecapability",
		"rhnpackagebreaks",
		"rhnpackagechangelogdata",
		"rhnpackagechangelogrec",
		"rhnpackageconflicts",
		"rhnpackageenhances",
		"rhnpackagefile",
		"rhnpackageobsoletes",
		"rhnpackagepredepends",
		"rhnpackageprovides",
		"rhnpackagerecommends",
		"rhnpackagerequires",
		"rhnsourcerpm",
		"rhnpackagesource",
		"rhnpackagesuggests",
	}
}

func DumpeChannelData(db *sql.DB, channelLabels []string) DataDumper {

	schemaMetadata := schemareader.ReadTablesSchema(db, readTableNames())

	initalDataSet := make([]processItem, 0)
	for _, channelLabel := range channelLabels {
		whereFilter := fmt.Sprintf("label = '%s'", channelLabel)
		sql := fmt.Sprintf(`SELECT * FROM rhnchannel where %s ;`, whereFilter)
		rows := executeQueryWithResults(db, sql)
		for _, row := range rows {
			initalDataSet = append(initalDataSet, processItem{schemaMetadata["rhnchannel"].Name, row, []string{"rhnchannel"}})
		}

	}
	tableData := dumpTableData(db, schemaMetadata, initalDataSet)
	PrintTableDataOrdered(db, schemaMetadata, tableData)
	return tableData
}
