package extract

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func TestExtract(t *testing.T) {
	reader := strings.NewReader(page)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
		return
	}

	result := extractFrom(doc)

	expected := []string{
		"HP DesignJet T520 36-in ePrinter Printer Accessories",
		"$3,427.00",
		"Brand: HP",
		"Manufacturer: HP",
		"CQ893A#B1K",
		"In the Box - HP Designjet T520 ePrinter, printhead, ink cartridges, printer stand, quick reference guide, setup poster, startup software, power cord, CD software, 1-year warrantyFeatures. Achieve accurate lines, sharp details up to 2400 dpi from the first print with a long-life printhead. Process complex files at high speeds, D prints in 35 seconds, thanks to 1 GB RAM and HP-GL/2 technology. The intuitive, full-color, 4.3-inch touchscreen simplifies navigation & printing. Easily set up your printer and connect everyone in your workgroup, thanks to built-in Wi-Fi",
		"Electronics",
		"Some stuff",
	}
	for _, value := range expected {
		if !strings.Contains(result, value) {
			t.Errorf("Should contain [%v]. No occurrence found in extraction result [%v]", value, result)
		}
	}
}

var page = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">

<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<body class=" catalog-product-view catalog-product-view product-hp-cq893a-b1k-designjet-t520-36-in-eprinter-printer-accessories categorypath-electronics-html category-electronics ">
<div id="root-wrapper">
<div class="wrapper">
        
    <div class="page">
         
        <div class="main-container col2-left-layout">
            <div class="main container show-bg">
                <div class="grid-full breadcrumbs">
    <ul>
                    <li class="home">
                            <a href="http://www.etronics.com/" title="Go to Home Page">Home</a>
                                        <span>/ </span>
                        </li>
                    <li class="category629">
                            <a href="http://www.etronics.com/electronics.html" title="">Electronics</a>
                                        <span>/ </span>
                        </li>
                    <li class="category999">
                            <a href="http://www.etronics.com/electronics.html" title="">Some stuff</a>
                                        <span>/ </span>
                        </li>
                    <li class="product">
                            <strong>HP DesignJet T520 36-in ePrinter Printer Accessories</strong>
                                    </li>
            </ul>
</div>
                
                <div class="col-main grid12-9 grid-col2-main in-col2">
                                        

<div class="product-view">

    <form action="http://www.etronics.com/checkout/cart/add/uenc/aHR0cDovL3d3dy5ldHJvbmljcy5jb20vaHAtY3E4OTNhLWIxay1kZXNpZ25qZXQtdDUyMC0zNi1pbi1lcHJpbnRlci1wcmludGVyLWFjY2Vzc29yaWVzLmh0bWw_X19fU0lEPVU,/product/36322/form_key/f5szcQSfNWMQ93Xf/" method="post" id="product_addtocart_form">


        <div class="product-shop grid12-5">

            <div class="product-name">
                <h1>HP DesignJet T520 36-in ePrinter Printer Accessories</h1>
            </div>
            
                        
                            <div class="short-description">
                    <div class="std">In the Box - HP Designjet T520 ePrinter, printhead, ink cartridges, printer stand, quick reference guide, setup poster, startup software, power cord, CD software, 1-year warrantyFeatures. Achieve accurate lines, sharp details up to 2400 dpi from the first print with a long-life printhead. Process complex files at high speeds, D prints in 35 seconds, thanks to 1 GB RAM and HP-GL/2 technology. The intuitive, full-color, 4.3-inch touchscreen simplifies navigation &amp; printing. Easily set up your printer and connect everyone in your workgroup, thanks to built-in Wi-Fi</div>
                </div>
            
                            <div class="sku"><span>SKU: </span>HPCQ893AB1K</div>
                            
            <div class="extrahint-wrapper">    <div class="product-pricing">
        HP DesignJet T520 36-in ePrinter Printer Accessories is available for purchase in increments of 1    </div>
</div>

<div class="product-type-data">
		    <p class="availability in-stock">Availability: <span>In stock</span></p>
	

                        
    <div class="price-box">
                                                                <span class="regular-price" id="product-price-36322">
                                            <span class="price">$3,427.00</span>                                    </span>
                        
        </div>

</div>
    
        </div> 
        
                     
        
        
    </form>
    
    
	<div class="box-additional box-tabs grid12-12">
    	<div id="product-tabs" class="gen-tabs gen-tabs-style1">

		        <ul class="tabs clearer">
							                	<li id="tab-description"><a href="http://www.etronics.com/hp-cq893a-b1k-designjet-t520-36-in-eprinter-printer-accessories.html#" class="">Description</a></li>
                            				            				                	<li id="tab-additional"><a href="http://www.etronics.com/hp-cq893a-b1k-designjet-t520-36-in-eprinter-printer-accessories.html#" class="current">Additional Info</a></li>
                            				            				                                </ul>
        <div class="tabs-panels"><h2 class="acctab" id="acctab-description">Description</h2><div class="panel" style="display: none;">    <h2>Details</h2>
    <div class="std">
        DesignJet T520 36-in ePrinter    </div>
</div><h2 class="acctab" id="acctab-additional">Additional Info</h2><div class="panel" style="display: block;">    <h2>Additional Info</h2>
    <table class="data-table" id="product-attribute-specs-table">
        <colgroup><col width="25%">
        <col>
        </colgroup><tbody>
                    <tr class="first odd">
                <th class="label">SKU</th>
                <td class="data last">HPCQ893AB1K</td>
            </tr>
                    <tr class="even">
                <th class="label">Manufacturer</th>
                <td class="data last">HP</td>
            </tr>
                    <tr class="odd">
                <th class="label">Model Number</th>
                <td class="data last">CQ893A#B1K</td>
            </tr>
                    <tr class="even">
                <th class="label">Length</th>
                <td class="data last">55.00</td>
            </tr>
                    <tr class="odd">
                <th class="label">Height</th>
                <td class="data last">23.00</td>
            </tr>
                    <tr class="last even">
                <th class="label">Width</th>
                <td class="data last">24.00</td>
            </tr>
                </tbody>
    </table>
</div></div>
        
	</div>
        
        	</div> 
    

        
</div> 


                </div>
                
                
            </div>
        </div>


    </div>
</div>
</div> 


</body></html>
`
