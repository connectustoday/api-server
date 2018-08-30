import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IOrganization extends IAccount {
	@prop({ required: true })
	preferred_name: string;

	@prop({ required: true })
	is_verified: boolean;

	@prop()
	opportunities?: Array<string>;

	@prop({ required: true })
	org_info: IOrganizationProfile;

	@prop()
	experience_validations?: Array<[string, string]>
}
